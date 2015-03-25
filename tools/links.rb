#!/usr/bin/env ruby

require 'colorize'
require 'net/http'
require 'nokogiri'
require 'uri'

# Hacky bit of Ruby to test URLs in the compiled service manual.
# Make sure to build it with Jekyll before running this.

LOG_LEVEL = 0 # 0 is everything, 1 is warnings, 2 is errors
MAKE_NETWORK_REQUESTS = false

html_files = Dir.glob('_site/service-manual/**/*.html')

visited_links = []

all_uris = []

$errors = false

def info(text)
  if LOG_LEVEL <= 0
    puts "[info] #{text}".light_black
  end
end

def warn(text)
  if LOG_LEVEL <= 1
    puts "[warning] #{text}".yellow
  end
end

def error(text)
  if LOG_LEVEL <= 2
    puts "[error] #{text}".red
  end
  $errors = true
end

def popular_domains(uris)
  uris.reject { |elem| elem.host.nil? }
      .group_by { |elem| elem.host }
      .map { |k, v| [k, v.count] }
      .sort_by { |k, v| -v }
end

html_files.each do |filename|
  f = File.open(filename)
  doc = Nokogiri::HTML(f)

  links = doc.css('a').map { |link| link['href'] }

  links.each do |link|

    # If a link is nil there's nothing more we can do with it, so eliminate it early
    if link.nil? or link == '#'
      warn("Skipping href with contents <#{link}> in #{filename}")
      next
    end

    # Test fragments before checking the visited links array, because fragments can
    # be repeated across pages
    if link.start_with?('#')
      if doc.css(link).length < 1
        warn("Fragment #{link} in #{filename} doesn't exist")
      elsif doc.css(link).length > 1
        warn("Fragment #{link} in #{filename} exists more than once")
      else
        next
      end
    end

    if visited_links.include? link
      next
    end

    uri = URI::parse(link)

    all_uris << uri

    if uri.scheme == 'mailto'
      info("No easy way to verify if a mailto link is valid <#{uri.to}>")
      next
    end

    if uri.host == 'www.gov.uk'
      warn("URI <#{uri}> contains production GOV.UK hostname rather than a relative path")
    end

    # Relative link to somewhere else on GOV.UK, so set the hostname manually
    if uri.host.nil?
      uri = URI::parse('https://www.gov.uk' + uri.path)
    end

    if uri.scheme != 'https'
      if uri.host.include? 'gov.uk'
        warn("Government URI <#{uri}> is not using HTTPS")
      else
        info("URI <#{uri}> is not using HTTPS")
      end
    end

    if uri.host == 'www.gov.uk' and uri.path.end_with?('/')
      warn("Government URI <#{uri}> will probably redirect to version without trailing slash")
    end

    if MAKE_NETWORK_REQUESTS
      http = Net::HTTP.new(uri.host, uri.port)

      if uri.scheme == 'https'
        http.use_ssl = true
      else
        http.use_ssl = false
      end

      request = Net::HTTP::Get.new(uri.request_uri)
      begin
        response = http.request(request)

        if response.code == '301' or response.code == '302'
          warn("#{response.code} response code from <#{uri}> to <#{response['location']}>")
        elsif response.code != '200'
          error("#{response.code} response code from <#{uri}>")
        end
      rescue StandardError => e
        error("#{e} trying to access <#{uri}>")
      end

      sleep 1
    end

    visited_links << link

  end

  f.close
end

# puts popular_domains(all_uris)

if $errors
  exit 1
end

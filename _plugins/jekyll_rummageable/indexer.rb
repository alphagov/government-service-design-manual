require 'rubygems'
require 'nokogiri'
require 'json'

module Jekyll

  class Indexer < Generator

    MANUAL_TITLE = "Government Service Design Manual"

    def initialize(config = {})
      super(config)

      @excludes = config['exclude_from_search'] || []
    end

    # Index all pages except pages matching any value in config['exclude_from_search']
    # The main content from each page is extracted and placed in a file for subsequent
    # indexing with rummager.
    def generate(site)
      puts 'Indexing pages...'

      # gather pages and posts
      items = site.pages + site.posts

      puts "Got details of #{items.length} items"

      items.reject! do |i|
        i.data['exclude_from_search'] ||
          page_should_be_excluded?(i) ||
          is_index_page?(i) ||
          ! i.output_is_html?
      end

      index = items.collect do |item|
        page_paragraphs = extract_text(site, item)

        {
          "_type"             => "manual_section",
          "organisations"     => ["government-digital-service"],
          "manual"            => "/service-manual",
          "title"             => "#{MANUAL_TITLE}: #{item.data['title'] || item.name}",
          "indexable_content" => page_paragraphs.join(" ").gsub("\r"," ").gsub("\n"," "),
          "description"       => extract_summary(site, item),
          "link"              => item.url
        }
      end
      index << {
        "_type"             => "manual",
        "organisations"     => ["government-digital-service"],
        "title"             => MANUAL_TITLE,
        "description"       => "All new digital services from the government must meet the Digital by Default Service Standard",
        "link"              => "/service-manual",
      }
      puts 'Indexing done, writing index to file'
      write_index_to_file(site, index)
      puts 'Index written to file'
    end

    protected
    # render the items, parse the output and get all text inside <p> elements
    def extract_text(site, page)
      doc = page_as_doc(site, page)
      doc.search('h1,h2,h3,h4,h5,li,blockquote,p').map { |e| e.text }
    end

    def extract_summary(site, page)
      doc = page_as_doc(site, page)
      doc.search('p').map { |e| e.text }.first
    end

    def page_should_be_excluded?(page)
      @excludes.any? { |s| (page.url =~ Regexp.new(s)) != nil }
    end

    def is_index_page?(page)
      page.is_a?(Jekyll::Page) && page.data['layout'].start_with?('role-index')
    end

    # We need DirectoryWatcher (the gem jekyll uses to detect changes to files
    # when in auto mode) to ignore the search index file, otherwise we hit an
    # infinite loop. Unfortunately that's not very configurable so using a
    # dotfile is the easiest way to achieve that.
    def write_index_to_file(site, index)
      File.open(".search-index.json", 'w') do |f|
        f.puts index.to_json
      end
    end

    private

    def page_as_doc(site, page)
      copy_of_page = page.dup
      copy_of_page.render({}, site.site_payload)
      Nokogiri::HTML(copy_of_page.output)
    end
  end
end

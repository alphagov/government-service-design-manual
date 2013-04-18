require 'rubygems'
require 'nokogiri'
require 'json'

module Jekyll

  class Indexer < Generator

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
      items = site.pages.dup.concat(site.posts)
      puts "Got details of #{items.length} items"

      items.reject! do |i|
        i.data['exclude_from_search'] ||
          page_should_be_excluded?(i) ||
          is_index_page?(i) ||
          ! i.output_is_html?
      end

      index = items.collect do |item|
        page_paragraphs = extract_text(site, item)

        puts 'Indexed ' << item.absolute_url

        {
          "title"             => item.data['title'] || item.name ,
          "indexable_content" => page_paragraphs.join(" ").gsub("\r"," ").gsub("\n"," "),
          "description"       => page_paragraphs.first,
          "link"              => item.absolute_url
        }
      end
      puts 'Indexing done'
      puts 'Writing index to file'
      write_index_to_file(site, index)
      puts 'Index written to file'
    end

    protected
    # render the items, parse the output and get all text inside <p> elements
    def extract_text(site, page)
      page.render({}, site.site_payload)
      doc = Nokogiri::HTML(page.output)
      doc.search('p').map {|e| e.text}
    end

    def page_should_be_excluded?(page)
      @excludes.any? { |s| (page.absolute_url =~ Regexp.new(s)) != nil }
    end

    def is_index_page?(page)
      page.is_a?(Jekyll::Page) && page.index?
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
  end 
end

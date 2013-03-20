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
    # The main content from each page is extracted and indexed at indextank.com
    # The doc_id of each indextank document will be the absolute url to the resource without domain name 
    def generate(site)
      puts 'Indexing pages...'
    
      # gather pages and posts
      items = site.pages.dup.concat(site.posts)

      # only process files that will be converted to .html and only non excluded files 
      items = items.find_all {|i| i.output_ext == '.html' && ! @excludes.any? {|s| (i.absolute_url =~ Regexp.new(s)) != nil } } 
      items.reject! {|i| i.data['exclude_from_search'] } 
      
      # dont process index pages
      items.reject! {|i| i.is_a?(Jekyll::Page) && i.index? }
			      
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
      
      File.open('search-index.json', 'w') do |f|
        f.puts index.to_json
      end
      
      puts 'Indexing done'
    end

    # render the items, parse the output and get all text inside <p> elements
    def extract_text(site, page)
      page.render({}, site.site_payload)
      doc = Nokogiri::HTML(page.output)
      doc.search('p').map {|e| e.text}
    end
  end 
end
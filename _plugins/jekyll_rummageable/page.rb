module Jekyll

  class Page
    # @origin https://github.com/kinnetica/jekyll-plugins/blob/master/sitemap_generator.rb
    attr_accessor :name

    def full_path_to_source
      File.join(@base, @dir, @name)
    end

    def output_is_html?
      output_ext == '.html'
    end
  end
end

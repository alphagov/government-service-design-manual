module Jekyll
  module TableOfContentsFilter
    def toc_list(document)
      toc = []
      content = Kramdown::Document.new(document)
      toc_tree = Kramdown::Converter::Toc.convert(content.root).first
      toc_tree.children.each do |heading|
        p heading
        html_id = heading.attr[:id]
        text = heading.value.options[:raw_text]
        toc << "<li class=\"level-#{html_id}\"><a href=\"\##{html_id}\">#{text}</a></li>"
      end
      "<ul>\n" + toc.join("\n") + "\n</ul>"
    end
  end
end

Liquid::Template.register_filter(Jekyll::TableOfContentsFilter)
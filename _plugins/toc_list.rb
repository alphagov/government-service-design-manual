module Jekyll
  module TableOfContentsFilter
    def toc_list(document)
      toc = []
      html = ''
      content = Kramdown::Document.new(document)
      toc_tree = Kramdown::Converter::Toc.convert(content.root).first
      html << "<ul>\n"
      toc_tree.children.each do |heading|
        html << show_tree(heading)
      end
      html << "</ul>\n"
    end
  end
end

def show_tree(my_item)
  html = ''
  html_id = my_item.attr[:id]
  text = my_item.value.options[:raw_text]
  html << "<li class=\"level-#{html_id}\"><a href=\"\##{html_id}\">#{text}</a>\n"

  if my_item.children.any?
    html << "<ul>\n"
    my_item.children.each do |my_child|
      html << show_tree(my_child)
    end
    html << "</ul>\n"
  end

  html << "</li>\n"

  return html
end

Liquid::Template.register_filter(Jekyll::TableOfContentsFilter)
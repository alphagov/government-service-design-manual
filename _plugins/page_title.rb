# encoding: UTF-8
module PageTitle
  class Generator < Jekyll::Generator
    def generate(site)
      page_title_suffix = 'Government Service Design Manual'
      site.pages.each do |page|
        page.data['title_without_suffix'] = page.data['title']
        if page.data['title']
          page.data['title'] = "#{page.data['title']} — #{page_title_suffix}"
        else
          page.data['title'] = page_title_suffix
        end
      end
    end
  end
end

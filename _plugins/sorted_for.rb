module Jekyll
  module SortedForImpl
    def render(context)
      sorted_collection = collection_to_sort context
      return if sorted_collection.empty?
      sort_attr = @attributes['sort_by']
      case_sensitive = @attributes['case_sensitive'] == 'true'
      i = sorted_collection.first
 
      if sort_attr != nil
        sorted_collection.reject!{|i| i.to_liquid[sort_attr].nil? }
        if i.to_liquid[sort_attr].instance_of? String and not case_sensitive
          sorted_collection.sort_by! { |i| i.to_liquid[sort_attr].downcase }
        else
          sorted_collection.sort_by! { |i| i.to_liquid[sort_attr] }
        end
      else
        sorted_collection.reject!(&:nil?)
        if i.instance_of? String and not case_sensitive
          sorted_collection.sort_by! { |i| i.downcase }
        else
          sorted_collection.sort!
        end
      end
 
      original_name = @collection_name
      result = nil
      context.stack do
        sorted_collection_name = "#{@collection_name}_sorted".sub('.', '_')
        context[sorted_collection_name] = sorted_collection
        @collection_name = sorted_collection_name
        result = super
        @collection_name = original_name
      end
      result
    end
  end
 
  class SortedForTag < Liquid::For
    include SortedForImpl
 
    def collection_to_sort(context)
      return context[@collection_name].dup
    end
 
    def end_tag
      'endsorted_for'
    end
  end
 
  class SortedKeysForTag < Liquid::For
    include SortedForImpl
 
    def collection_to_sort(context)
      return context[@collection_name].keys
    end
 
    def end_tag
      'endsorted_keys_for'
    end
  end
end
 
Liquid::Template.register_tag('sorted_for', Jekyll::SortedForTag)
Liquid::Template.register_tag('sorted_keys_for', Jekyll::SortedKeysForTag)
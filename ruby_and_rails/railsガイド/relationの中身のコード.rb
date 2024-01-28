# rails ソースコード
# activerecord/lib/active_record/associations.rb
# ===== =====

# frozen_string_literal: true

module ActiveRecord
  # See ActiveRecord::Associations::ClassMethods for documentation.
  # https://api.rubyonrails.org/classes/ActiveRecord/Associations/ClassMethods.html
  module Associations # :nodoc:
    extend ActiveSupport::Autoload
    extend ActiveSupport::Concern

    # These classes will be loaded when associations are created.
    # So there is no need to eager load them.
    autoload :Association
    autoload :SingularAssociation
    autoload :CollectionAssociation
    autoload :ForeignAssociation
    autoload :CollectionProxy
    autoload :ThroughAssociation

    module Builder # :nodoc:
      autoload :Association,           "active_record/associations/builder/association"
      autoload :SingularAssociation,   "active_record/associations/builder/singular_association"
      autoload :CollectionAssociation, "active_record/associations/builder/collection_association"

      autoload :BelongsTo,           "active_record/associations/builder/belongs_to"
      autoload :HasOne,              "active_record/associations/builder/has_one"
      autoload :HasMany,             "active_record/associations/builder/has_many"
      autoload :HasAndBelongsToMany, "active_record/associations/builder/has_and_belongs_to_many"
    end

    eager_autoload do
      autoload :BelongsToAssociation
      autoload :BelongsToPolymorphicAssociation
      autoload :HasManyAssociation
      autoload :HasManyThroughAssociation
      autoload :HasOneAssociation
      autoload :HasOneThroughAssociation

      autoload :Preloader
      autoload :JoinDependency
      autoload :AssociationScope
      autoload :DisableJoinsAssociationScope
      autoload :AliasTracker
    end

    def self.eager_load!
      super
      Preloader.eager_load!
      JoinDependency.eager_load!
    end

    # Returns the association instance for the given name, instantiating it if it doesn't already exist
    def association(name) # :nodoc:
      association = association_instance_get(name)

      if association.nil?
        unless reflection = self.class._reflect_on_association(name)
          raise AssociationNotFoundError.new(self, name)
        end
        association = reflection.association_class.new(self, reflection)
        association_instance_set(name, association)
      end

      association
    end

    def association_cached?(name) # :nodoc:
      @association_cache.key?(name)
    end

    def initialize_dup(*) # :nodoc:
      @association_cache = {}
      super
    end

    private
      def init_internals
        super
        @association_cache = {}
      end

      # Returns the specified association instance if it exists, +nil+ otherwise.
      def association_instance_get(name)
        @association_cache[name]
      end

      # Set the specified association instance.
      def association_instance_set(name, association)
        @association_cache[name] = association
      end

      # more complex than the simple and guessable ones possible.
      module ClassMethods

        def has_many(name, scope = nil, **options, &extension)
          reflection = Builder::HasMany.build(self, name, scope, options, &extension)
          Reflection.add_reflection self, name, reflection
        end


        def has_one(name, scope = nil, **options)
          reflection = Builder::HasOne.build(self, name, scope, options)
          Reflection.add_reflection self, name, reflection
        end


        def belongs_to(name, scope = nil, **options)
          reflection = Builder::BelongsTo.build(self, name, scope, options)
          Reflection.add_reflection self, name, reflection
        end
      end
  end
end

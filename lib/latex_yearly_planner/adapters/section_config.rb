# frozen_string_literal: true

module LatexYearlyPlanner
  module Adapters
    class SectionConfig
      attr_reader :name, :section_config, :planner_config

      def initialize(name:, section_config:, planner_config:)
        @name = name
        @section_config = section_config
        @planner_config = planner_config
      end

      def enabled?
        section_config.enabled
      end

      def instantiate
        constant.new(name:, header:, body:, section_config:, planner_config:)
      end

      private

      def header
        header_constant.new(name:, section_config:, planner_config:)
      end

      def body
        body_constant.new(name:, section_config:, planner_config:)
      end

      def constant
        "LatexYearlyPlanner::Planners::#{template}::Sections::#{camelized_name}".constantize
      end

      def header_constant
        "LatexYearlyPlanner::Planners::#{template}::Components::#{camelized_name}Header".constantize
      end

      def body_constant
        "LatexYearlyPlanner::Planners::#{template}::Components::#{camelized_name}Body".constantize
      end

      def template
        planner_config.template.camelize
      end

      def camelized_name
        name.to_s.camelize
      end
    end
  end
end

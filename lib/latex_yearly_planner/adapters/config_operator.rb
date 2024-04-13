# frozen_string_literal: true

module LatexYearlyPlanner
  module Adapters
    class ConfigOperator
      attr_reader :section_config, :planner_config

      def initialize(section_config:)
        @section_config = section_config
        @planner_config = section_config.planner_config
      end

      def get(*keys)
        section_config.section_config.dig(:parameters, *keys) ||
          planner_config.planner.dig(:parameters, *keys)
      end

      def object(key)
        section_config.object(key) || planner_config.object(key)
      end

      def placement(key)
        section_config.placement(key) || planner_config.placement(key)
      end

      def months
        (start_date..end_date)
          .select { |date| date.mday == 1 }
          .map(&method(:initialize_month))
      end

      def quarters
        months.each_slice(3).map(&:first).map(&:quarter)
      end

      private

      def start_date
        @start_date ||= Date.parse(get(:start_date))
      end

      def end_date
        @end_date ||= Date.parse(get(:end_date))
      end

      def initialize_month(date)
        Calendar::Month.new(weekday_start:, year: date.year, month: date.month)
      end

      def weekday_start
        @weekday_start ||= get(:weekday_start).downcase.to_sym
      end
    end
  end
end
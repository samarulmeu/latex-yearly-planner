# frozen_string_literal: true

module LatexYearlyPlanner
  module TeX
    class VerticalPhantom
      def initialize(content)
        @content = content
      end

      def to_s
        "\\vphantom{#{content}}"
      end
    end
  end
end
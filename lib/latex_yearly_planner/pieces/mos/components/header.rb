# frozen_string_literal: true

module LatexYearlyPlanner
  module Pieces
    module Mos
      module Components
        class Header < Component
          def top_table
            table = TeX::Tabular.new(**parameters(:header, :top_table))
            table.add_row([calendar_cell])

            table
          end

          def calendar_cell
            cell = TeX::Cell.new('Calendar')
            cell.selected = true if section_name == :annual

            cell
          end

          def margin_note(highlight_quarters: [], highlight_months: [])
            <<~LATEX
              \\marginnote{%
                #{quarter_table_from_months(highlight_quarters:)}
                \\\\[#{param(:header, :between_tables_spacing)}]
                #{table_from_months(highlight_months:)}
              }
            LATEX
              .strip
          end

          def quarter_table_from_months(highlight_quarters:)
            table = TeX::TabularX.new(**parameters(:header, :quarterly_table))

            all_quarters.map do |quarter|
              cell = TeX::Cell.new(quarter.name)
              cell.selected = true if highlight_quarters.include?(quarter)

              table.add_row([cell])
            end

            table
          end

          def table_from_months(highlight_months:)
            table = TeX::TabularX.new(**parameters(:header, :monthly_table))

            all_months.each do |month|
              cell = TeX::Cell.new(month.name[0..2])
              cell.selected = true if highlight_months.include?(month)

              table.add_row([cell])
            end

            table
          end
        end
      end
    end
  end
end
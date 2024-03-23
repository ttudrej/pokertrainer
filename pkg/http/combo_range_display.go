package http

import (
	"fmt"
	"strconv"

	"github.com/ttudrej/pokertrainer/debugging"
)

/*
#########################################################################
#########################################################################
#########################################################################

######## ##     ## ##    ##  ######   ######
##       ##     ## ###   ## ##    ## ##    ##
##       ##     ## ####  ## ##       ##
######   ##     ## ## ## ## ##        ######
##       ##     ## ##  #### ##             ##
##       ##     ## ##   ### ##    ## ##    ##
##        #######  ##    ##  ######   ######

#########################################################################
#########################################################################
#########################################################################
*/

// Based on python code from here:
// https://poker.readthedocs.io/en/latest/_modules/poker/hand.html#Range.to_html

// #################################################################
// drawComleteRange_HTML create an HTML table representing a full "range chart",
// a 13x13 grid for use in range representation.
func drawCompleteRangeHTML() string {
	Info.Println(debugging.ThisFunc())
	var html string
	// html = `<table class="range">`
	html = `<table>`

	var shape string
	// var cssClass string

	var twoCardHand string

	for row := 13; row >= 1; row-- {
		html += `<tr>`
		for col := 13; col >= 1; col-- {

			/*
				if row > col {
					shape, cssClass = "s", "suited"
				} else if row < col {
					shape, cssClass = "o", "offsuit"
				} else {
					shape, cssClass = "p", "pair"
				}
			*/

			// html += Info.Sprintf(`<td class="%s">`, cssClass)
			html += fmt.Sprintf(`<td>`)
			twoCardHand = strconv.Itoa(row) + strconv.Itoa(col) + shape
			html += twoCardHand
			html += `</td>`
		}
		html += `</tr>`
	}
	html += `</table>`

	return html
}

// #################################################################

// #################################################################

/*
def to_html(self):
"""Returns a 13x13 HTML table representing the range.

The table's CSS class is ``range``, pair cells (td element) are ``pair``, offsuit hands are
``offsuit`` and suited hand cells has ``suited`` css class.
The HTML contains no extra whitespace at all.
Calculating it should not take more than 30ms (which takes calculating a 100% range).
"""

# note about speed: I tried with functools.lru_cache, and the initial call was 3-4x slower
# than without it, and the need for calling this will usually be once, so no need to cache

html = ['<table class="range">']

for row in reversed(Rank):
	html.append('<tr>')

	for col in reversed(Rank):
		if row > col:
			suit, cssclass = 's', 'suited'
		elif row < col:
			suit, cssclass = 'o', 'offsuit'
		else:
			suit, cssclass = '', 'pair'

		html.append('<td class="%s">' % cssclass)
		hand = Hand(row.val + col.val + suit)

		if hand in self.hands:
			html.append(unicode(hand))

		html.append('</td>')

	html.append('</tr>')

html.append('</table>')
return ''.join(html)
*/

/*
// #################################################################
[docs]    def to_ascii(self, border=False):
"""Returns a nicely formatted ASCII table with optional borders."""

table = []

if border:
	table.append('┌' + '─────┬' * 12 + '─────┐\n')
	line = '├' + '─────┼' * 12 + '─────┤\n'
	border = '│ '
	lastline = '\n└' + '─────┴' * 12 + '─────┘'
else:
	line = border = lastline = ''

for row in reversed(Rank):
	for col in reversed(Rank):
		if row > col:
			suit = 's'
		elif row < col:
			suit = 'o'
		else:
			suit = ''

		hand = Hand(row.val + col.val + suit)
		hand = unicode(hand) if hand in self.hands else ''
		table.append(border)
		table.append(hand.ljust(4))

	if row.val != '2':
		table.append(border)
		table.append('\n')
		table.append(line)

table.append(border)
table.append(lastline)

return ''.join(table)
*/

/*
CSS defs:

<style>
   td {
      //  Make cells same width and height and centered //
      width: 30px;
      height: 30px;
      text-align: center;
      vertical-align: middle;
   }
   td.pair {
      background: #aaff9f;
   }
   td.offsuit {
      background: #bbced3;
   }
   td.suited {
      background: #e37f7d;
   }
</style>
*/

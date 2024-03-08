package gameobjects

type gameName string
type gameType string

// gameSize expressed in cents ($0.01). A 1c/2c is "NL 2", a 50c/$100 is "NL 100", a $1/$2 is "NL 200", ...
type gameSize struct {
	blind1 int // small blind in $0.01
	blind2 int // big blind
}

type game struct {
	gName gameName
	gType gameType
	gSize gameSize
}

const (
	nl2   gameName = "NL 1c/2c"
	nl4   gameName = "NL 2c/4c"
	nl200 gameName = "NL $1/$2"
	nl300 gameName = "NL $1/$3"
	nl500 gameName = "NL $2/$5"

	nlh gameType = "NLH"
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
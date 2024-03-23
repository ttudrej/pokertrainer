// Package hand_rcording documentation
// Represents all the functionality of a system that can store and recall/replay any/all hands
// played anywhere in the poker room, all/any tables.
// All the methods and functions required for storing/recording the progress of hands on any table.
// We probably want to choose a popular standard for the format.
// Whatever the solution, the goal is to have a systyem in place that will allow
// "replaying" a hand. A replay should show all the analysis one would see during a live
// hand, and potentially deeper levels of analysis that my not be possible during live play
// (I don't know what those would be atm)
// Ownership of records may be a consideration here at some poit, ie. who can look at what hands.
// Can player A look at player B's hands, if A was not at the same table at the same time,
// for example.
package hand_recording

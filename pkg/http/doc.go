// Package http documentation
// All functinality required for rendering a table and a corresponding
// game in a browser, from a point of view of an individual player (to start with), with
// all the control elements, bet/call/fold buttons, bet size sliders, etc.
// The "state" of the game is fully managed by the back end.
// Here, we're responsible for taking player inputs, transmitting them to the back end, and
// updating the displayed table state after all actions, originated by this or
// any other player or dealer.
package http

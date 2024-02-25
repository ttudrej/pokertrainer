
# pokertrainer

## Desription

Provides a training environment for playing poker hands. A simulation which calculates and displays various statistics based on the state of the hand and action so far.

## Detailed Descritpion

A poker game simulator which also calculates and shows the statistics that should be taken
into consideration when making decisions. Allows implementation of various strategies.  
The intent is to give the user all the information that can be gleaned with use of statistics,
combined with player styles and various strategies they may be using in specific scenarios. This
allows the user to compare their assumptions and calculations to those derived by the applicaiton for
immediate feedback.

## TableVision Feed Mode

The game itself, instead of being randomly simulated, can be driven by a real/online game feed via an HTTP API.

## Dependencies

### Input Viewer
Something nees to be feeding our API endpoint with table state data.

## Project Directory Layout
Idea for repo layout came from [Hexagonal Architecture](https://youtu.be/oL6JBUk6tj0?t=1613) model.

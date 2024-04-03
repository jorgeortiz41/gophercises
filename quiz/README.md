# Quiz Game
Implemented a quiz game that reads questions and answers from a CSV file and poses them to the user. The game keeps track of the number of correct answers and at the end, it displays the final score.


## Part 1
Created a program that will read in a quiz provided via a CSV file (more details below) and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect.

The CSV file will be in a format like below:
``` 
5+5,10
1+2,3
8+3,11
```

At the end of the quiz the program outputs the total number of questions correct and how many questions there were in total. Questions given invalid answers are considered incorrect.

## Part 2
Added a timer to the quiz. The quiz should be timed and if the user doesn't provide an answer in the allotted time then it should be considered incorrect. The default time limit is  30 seconds. The user is be able to customize the time limit.
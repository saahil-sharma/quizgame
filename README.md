## QuizGame

Basically, this program is a Go quiz that I have made using Go, which takes in your answers and responses as a csv file. 
There is a timer to control the amount of time that the user spends on the quiz. The default value is 30 seconds. 

The csv file should be formatted with rows of two columns, with the first column including the question and the second column with the answer.
See the example file "problems.csv" for questions. 

There are two flags included with this CLI application, the first being an alternate name for the CSV file. 
The default name is 'problems.csv', but it can be changed with a flag named -csv="alternateName.csv". 
The second flag is for the timer. This can be changed by setting a second (or first) flag, of -limit="20" for example, changing the time limit for the quiz to 20 seconds. 

To use the quiz game yourself, you should clone the repository (paste this in your terminal):
```
git clone https://github.com/saahil-sharma/quizgame
```
Change your directory to quizgame:
```
cd quizgame
```
Build the project with the command:
```
go build .
```
This should create an executable of the name 'quizgame' which can be ran with:
```
./quizgame
```
This does require that you have the Go language installed here, which can be found [here](https://go.dev). 

Thanks for checking out my code! Chances are there might be something wrong or not working, and I would love any feedback. 
You can either submit a PR or email me at "saahilsharma1017[@]gmail[dot]com". 

Have a good day. 

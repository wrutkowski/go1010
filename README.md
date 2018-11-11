# go1010

![go1010 game play](https://raw.githubusercontent.com/wrutkowski/go1010/master/assets/go1010.gif)

go1010 is a game running in terminal developed using Go. The game consists of a 10x10 size game board and three blocks with random figures. Player has to place the block in the empty places on the game board for which they receive points. When there is a full row or column on the game board it gets deleted and player receives points. Game is over when there is no possiblity to place any of the available block on the board.

Game is inspired by a mobile game called `1010!`.

### AI

Neural Network was created to play go1010. There are several command to oversee the training and save / load the neural networks. After running the training for around 2 hours on population of 500 and reaching generation ~4400 fitness of 37 was reached. Fitness is directly tied to game score.

![go1010 AI playing](https://raw.githubusercontent.com/wrutkowski/go1010/master/assets/ai_playing_f37.gif)

This is the instruction set available:

```
Enter - next iteration
s NUM - skip NUM of steps
g NUM - skip NUM generations
f NUM - until Fitness is above NUM
t NUM - run for NUM seconds
ng - run until next generation
drawing on/off - enable/disable drawing each iteration
save filename - saves top performant Neural Network to a file
load filename - loads Neural Network to last place
help - this help
e - exit 
```

Also, I think at one point the neural network wanted to tell me something ;-)

![go1010 AT telling something F*](https://raw.githubusercontent.com/wrutkowski/go1010/master/assets/game_f.png)

### License

Code is released under MIT License (MIT)

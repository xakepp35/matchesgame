import React from 'react';
import logo from './logo.jpg';
import './App.css';
import {NavLink, Switch, Route} from 'react-router-dom';
// import ReactBootstrapSlider from 'react-bootstrap-slider';

import { Container, InputGroup, InputGroupAddon, InputGroupText, Input, Button } from 'reactstrap';

import RestApi from "./RestApi";

function LandingView() {
  return(
    <div>
      Landing Пага. Тут может быть ваша реклама :)
    </div>
  )
}

const NewGameView = (props) => {
  return(
    <Container className="App-cont">

      <h2 className="App-padded">
        Начать новую партию
      </h2>

      <InputGroup className="py-2">
        <InputGroupAddon addonType="prepend">
          <InputGroupText>Player Name:</InputGroupText>
        </InputGroupAddon>
        <Input placeholder="mr.Junior" name="PlayerName" value={props.PlayerName} onChange={props.OnChange} />
      </InputGroup>

      <InputGroup className="py-2">
        <InputGroupAddon addonType="prepend">
          <InputGroupText>Начальное количество спичек:</InputGroupText>
        </InputGroupAddon>
        <Input placeholder="30" name="StartMatchesAmount" value={props.StartMatchesAmount} onChange={props.OnChange} />
      </InputGroup>

      <InputGroup className="py-2">
        <InputGroupAddon addonType="prepend">
          <InputGroupText>Макс кол-во спичек за ход:</InputGroupText>
        </InputGroupAddon>
        <Input placeholder="3" name="MaxMatchesPrerTurn" value={props.MaxMatchesPrerTurn} onChange={props.OnChange} />
      </InputGroup>

      <Button onClick={props.NewGameHandler}>Go!</Button>

    </Container>
  )
}

class NewGameApp extends React.Component {

  constructor(props) {
    super(props);
    this.state= {
      PlayerName: "mr.Junior",
      MaxMatchesPrerTurn: 3,
      StartMatchesAmount: 30,
    }
    this.handleChange = this.handleChange.bind(this)
    this.startNewGame = this.startNewGame.bind(this)
  }

  handleChange(event) {
    this.setState({[event.target.name]: event.target.value});
  }

  async startNewGame() {
    //console.dir(this.state);
    const res = await RestApi.startNewGame({
      PlayerName: this.state.PlayerName,
      MaxMatchesPrerTurn: this.state.MaxMatchesPrerTurn,
      StartMatchesAmount: this.state.StartMatchesAmount,
    })
    console.dir(res)
    //navigate to /game
  }
  

  render() {
    return <NewGameView OnChange={this.handleChange} NewGameHandler={this.startNewGame} {...this.state}/>
  }
}

function TopScoresView() {
  return(
    <div>
      TopScores
    </div>
  )
}

function GameSessionView() {
  return(
    <div>
      GameSession
    </div>
  )
}

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <NavLink
          className="App-link"
          to="/new"
          rel="noopener noreferrer"
        >
          New Game
        </NavLink>
        <NavLink
          className="App-link"
          to="/scores"
          rel="noopener noreferrer"
        >
          Top Scores
        </NavLink>
        <a
          className="App-link"
          href="https://github.com/xakepp35/matchesGame"
          target="_blank"
          rel="noopener noreferrer"
        >
          Github
        </a>
      </header>
      <main className="App-main">
      <Switch>
        <Route exact path='/' component={LandingView}/>
        <Route path='/new' component={NewGameApp}/>
        <Route path='/scores' component={TopScoresView}/>
        <Route path='/game' component={GameSessionView}/>
        </Switch>
      </main>
    </div>
  );
}

export default App;

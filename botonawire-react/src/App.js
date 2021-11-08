import React from "react";
import "./App.css"

class RobotHeader extends React.Component {
  constructor(props) {
    super(props);

    this.connectedColor = this.props.robot.connected ? "green" : "red"
    this.connectedString = this.props.robot.connected ? "Connected" : "Disconnected"

    this.onClick = this.onClick.bind(this)
  }

  onClick() {
    this.props.switchRobot(this.props.robot)
  }

  render() {
    return (
      <div class="robotHeader" onClick={this.onClick}>
        {this.props.robot.name} - <text style={{color: this.connectedColor}}>{this.connectedString}</text>
      </div>
    )
  }
}

class StatusBar extends React.Component {
  constructor(props) {
    super(props);
    this.state = { robots: [] }
  }

  componentDidMount() {
    fetch('/robots').then(res => res.json()).then(data => {
      if(data != null) {
        this.setState({ robots: data })
      }
    })

    this.interval = setInterval(() => (fetch('/robots').then(res => res.json()).then(data => {
      if(data != null) {
        this.setState({ robots: data })
      }
    })), 10000)
  }

  render() {
    return (
      <div class="sidebar">
        {this.state.robots.map((robot) => {
          return <RobotHeader robot={robot} switchRobot={this.props.switchRobot}/>
        })}
      </div>
    )
  }
}

class RobotControllerForward extends React.Component {
  constructor(props) {
    super(props);

    this.enabledColor = this.props.robot.automode ? "green" : "red"
    this.enabledString = this.props.robot.automode ? "Forward" : "Forward"

    this.onClick = this.onClick.bind(this)
  }

  onClick() {
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(
        { 
          "uuid": this.props.robot.uuid,
        })
    };
    fetch("/forward", requestOptions)
  }

  render() {
    return (
      <div class="robotButton" onClick={this.onClick}>
        <text style={{color: this.enabledColor}}>{this.enabledString}</text>
      </div>
    )
  }
}

class RobotControllerBackward extends React.Component {
  constructor(props) {
    super(props);

    this.enabledColor = this.props.robot.automode ? "green" : "red"
    this.enabledString = this.props.robot.automode ? "Backwards" : "Backwards"

    this.onClick = this.onClick.bind(this)
  }

  onClick() {
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(
        { 
          "uuid": this.props.robot.uuid,
        })
    };
    fetch("/backward", requestOptions)
  }

  render() {
    return (
      <div class="robotButton" onClick={this.onClick}>
        <text style={{color: this.enabledColor}}>{this.enabledString}</text>
      </div>
    )
  }
}

class RobotControllerStop extends React.Component {
  constructor(props) {
    super(props);

    this.enabledColor = this.props.robot.automode ? "green" : "red"
    this.enabledString = this.props.robot.automode ? "Stop" : "Stop"

    this.onClick = this.onClick.bind(this)
  }

  onClick() {
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(
        { 
          "uuid": this.props.robot.uuid,
        })
    };
    fetch("/stop", requestOptions)
  }

  render() {
    return (
      <div class="robotButton" onClick={this.onClick}>
        <text style={{color: this.enabledColor}}>{this.enabledString}</text>
      </div>
    )
  }
}
class Control extends React.Component {

  render() {
    return (
      <div class="control">
        Robot Name: {this.props.robot.name}<br />
        Charge: [{this.props.robot.charge}]<br />
        UUID: {this.props.robot.uuid}<br />
        <RobotControllerForward robot={this.props.robot} /><br />
        <RobotControllerBackward robot={this.props.robot} /><br />
        <RobotControllerStop robot={this.props.robot} />
      </div>
    )
  }
}

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = { currentRobotParams: {} }
    this.switchRobot = this.switchRobot.bind(this)
  }

  switchRobot(robot) {
      this.setState({ currentRobotParams: robot })
  }

  render() {
    return (
      <div class="background-container">
        <StatusBar switchRobot={this.switchRobot}/>
        <Control robot={this.state.currentRobotParams} />
      </div>
    );
  }
}

export default App;

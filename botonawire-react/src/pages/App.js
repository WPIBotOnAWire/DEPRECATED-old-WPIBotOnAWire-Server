import React from "react";
import "../style/App.css"
import "../components/Header"
import "./Robot"
import Header from "../components/Header";
import RobotStatus from "./Robot";

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = { currentRobotParams: {} }
    this.switchRobot = this.switchRobot.bind(this)
  }

  switchRobot(robot) {
      this.setState({ currentRobotParams: robot })
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
      <div class="background-container">
        <Header />
        <RobotStatus robot={this.props.robots[0]}/>
      </div>
    );
  }
}

export default App;
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

  render() {
    return (
      <div class="background-container">
        <Header />
        <RobotStatus />
      </div>
    );
  }
}

export default App;
import React from "react";
import "../components/StatusCard"
import StatusCard from "../components/StatusCard";
import MapCard from "../components/MapCard";
import "../style/App.css"

class RobotStatus extends React.Component {

  render() {
    return (
        <div class="cardview">
            <StatusCard />
            <MapCard />
        </div>
    )
  }
}

export default RobotStatus;
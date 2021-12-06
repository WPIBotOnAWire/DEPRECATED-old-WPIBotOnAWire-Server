import React, { Component } from "react";
import { Card } from "react-bootstrap";
import "../style/App.css"

class StatusCard extends Component {
    constructor(props) {
        super(props)
        this.props.robot = props.robot
        this.led_toggle = this.led_toggle.bind(this)
    }

    led_toggle() {
        const requestOptions = {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(
            { 
              "uuid": this.props.robot.uuid,
            })
        };
        fetch("/led", requestOptions)
    }

    render() {
        return(
            <div class="card">
                Robot: {this.props.robot.uuid}
                <button onClick={this.led_toggle}>THIS IS A BUTTON</button>
            </div>
        )
    }
}

export default StatusCard;
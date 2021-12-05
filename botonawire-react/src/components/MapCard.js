import React, { Component } from "react";
import "../style/App.css";
import MapContainer from "./MapContainer";

class MapCard extends Component {
    render() {
        return(
            <div class="card">
                <MapContainer />
            </div>
        )
    }
}

export default MapCard;
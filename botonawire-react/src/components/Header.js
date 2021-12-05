import React from "react";
import "../style/Header.css"
import logo from "../assets/logo_color.png"
import { FaBars } from "react-icons/fa"
import { CgProfile } from "react-icons/cg"

class Header extends React.Component {
    render() {
        return(
            <div class="header">
                <div class="left">
                    <FaBars class="bars" size={30}/>
                </div>

                <div class="center">
                    <img src={logo} class="logo"/>
                </div>

                <div class="right">
                    <CgProfile class="profile" size={50}/>
                </div>
            </div>
        )
    }
}

export default Header
/** @jsx React.DOM */
var React = require('react');
var ReactDOM = require('react-dom');

class Hello extends React.Component{
    render(){
        let who = "World react";
        return who
    }
}

ReactDOM.render(<Hello/>, document.querySelector("#root"));
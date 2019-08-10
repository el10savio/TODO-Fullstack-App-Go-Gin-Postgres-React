import React from "react";
import "./styles/Todolist.css";

class Todolist extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      items: [],
    };
  }

  componentDidMount() {
    fetch("http://localhost:8081/items")
      .then(res => res.json())
      .then(json => console.log(json));
  }

  render() {
    return (
      <div className="TodoList">
        <div className="List">
          <div className="ListItem">
            <div className="Title">Title</div>
            <div className="Status">Status</div>
          </div>
          <div className="ListItem">
            <div className="Title">Title</div>
            <div className="Status">Status</div>
          </div>
          <div className="ListItem">
            <div className="Title">Title</div>
            <div className="Status">Status</div>
          </div>
        </div>
      </div>
    );
  }
}

export default Todolist;

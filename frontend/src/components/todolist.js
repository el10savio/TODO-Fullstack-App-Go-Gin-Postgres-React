import React from "react";
import "./styles/Todolist.css";

class Todolist extends React.Component {
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

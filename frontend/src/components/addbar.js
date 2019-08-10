import React from "react";
import "./styles/Addbar.css";

class AddBar extends React.Component {
  render() {
    return (
      <div className="AddBar">
        <input
          className="AddBar-Text"
          type="text"
          placeholder="Enter TODO Item"
        />
        <div className="AddBar-Button">Add Item</div>
      </div>
    );
  }
}

export default AddBar;

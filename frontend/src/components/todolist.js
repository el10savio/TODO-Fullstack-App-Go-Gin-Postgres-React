import React from "react";
import "./styles/Todolist.css";

class Todolist extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      items: [],
    };
  }

  createItem(item) {
    return (
      <div className="ListItem" key={item.id} id={item.id}>
        <div className="Title">{item.item}</div>
        <div className="Status">{item.done && "Done"}</div>
      </div>
    );
  }

  componentDidMount() {
    fetch("http://localhost:8081/items")
      .then(res => res.json())
      .then(json => this.setState({ items: json.items }));
  }

  render() {
    var items = this.state.items;
    return (
      <div className="TodoList">
        <div className="List">{items.map(item => this.createItem(item))}</div>
      </div>
    );
  }
}

export default Todolist;

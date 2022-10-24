import React, { Component } from 'react';
import './App.css';

function Title(props) {
  return (
    <div class="row">
      <div class="col">
        <h1>{props.title}</h1>
      </div>
    </div>
  )
}

async function getTableNames() {
  const response = await fetch("http://localhost:8080/get_table_names", {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    }
  })
  const data = await response.json()
  return data.tableNames
}

async function getList(listName) {
  const response = await fetch(`http://localhost:8080/get_table?tableName=${listName}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    }
  })
  return await response.json()
}

class App extends Component {
  constructor(props) {
    super(props)
    this.state = { tableNames: [] }
  }

  async componentDidMount() {
    try {
      let names = await getTableNames()
      this.setState({ tableNames: names, selectedList: "" })
    } catch (e) {
      console.log(e)
    }
  }

  setCurrentList = (newList) => {
    console.log(`setting current list to ${newList}`)
    this.setState({ selectedList: newList })
  }

  render() {
    return (
      <div class="container text-center" className="App">
        <Title title="This is my appy poo poo" />
        <div class="row">
          <div class="col">
            <ListSelector tableNames={this.state.tableNames} setCurrentList={this.setCurrentList}/>
            <TodoList />
          </div>
        </div>
      </div>
    );
  }
}

class Entry extends Component {
  constructor(props) {
    super(props)
    this.props = props
  }

  render() {
    return (
      <div>
        <p>Entry: {this.props.summary}</p>
      </div>
    )
  }
}

class ListSelector extends Component {
  constructor(props) {
    super(props)
    this.props = props
  }

  handleSelect = (event) => {
    this.props.setCurrentList(event.target.value)
  }

  render() {
    let listItems = this.props.tableNames.map((name) => <option value={name}>{name}</option>)
    return (
      <div class="list-selector">
        <select onChange={this.handleSelect}>
          <option value="">Select a list</option>
          {listItems}
        </select>
      </div>
    )
  }
}

class TodoList extends Component {
  constructor(props) {
    super(props);
    this.state = {
      items: [
        "Hello fuego",
        "Dello buego",
      ]
    };
  }

  async componentDidMount() {

  }

  render() {
    return (
      <div>
        <ul>
          {this.state.items.map((sum) => <li><Entry summary={sum} /></li>)}
        </ul>
      </div>
    );
  }
}

export default App;

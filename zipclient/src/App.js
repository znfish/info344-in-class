import React from "react";
import $ from "jquery";


var App = React.createClass({
    getInitialState: function () {
        return { url: '', res: '' };
    },

    getName: function () {
        console.log('clicked')
        var name = $("#name").val();
        console.log(name)
        var url1 = "http://localhost:4000/hello?name=" + name;
        this.getResponse(url1);
    },

    getResponse: function (url) {
        var self = this;
        console.log(url)
        $.get(url, function (data) {
            self.setState({ res: data })
            console.log(self.state.res)
        });
    },

    render: function () {
        return (
            <div className="App">
                <h1>Test</h1>
                <form>
                    <input type="text" name="name" id="name" />
                </form>
                <button id="button" onClick={this.getName}>Submit</button>
                {
                    this.state.res !== '' &&
                    <h4>{this.state.res}</h4>
                }
            </div>
        );
    }

});

export default App;
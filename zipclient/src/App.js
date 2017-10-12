import React from "react";
import $ from "jquery";


var App = React.createClass({
    getInitialState: function () {
        return { url: '', res: [] };
    },

    getCity: function () {
        console.log('clicked')
        var city = $("#city").val();
        console.log(city)
        var url1 = "http://localhost:4000/zips/" + city;
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
                <h1>City</h1>
                <form>
                    <input type="text" id="city" />
                </form>
                <button id="button" onClick={this.getCity}>Submit</button>
                {
                    this.state.res !== '' &&
                    this.state.res.map(function (d, i) {
                        return (
                            <div key={i}>
                                <h4>{d.Code} {d.City} {d.State}</h4>
                            </div>
                        );
                    })
                }
            </div>
        );
    }

});

export default App;
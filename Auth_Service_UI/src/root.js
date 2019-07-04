import React, {Component} from 'react';
import TopTab from "./tab/tab";

class Root extends Component {

    render() {
        return (
            <div>
                <div>
                    <TopTab/>
                </div>
                <div>
                    {this.props.children}
                </div>
            </div>

        );
    }
}

export default (Root);
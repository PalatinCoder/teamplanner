import { LitElement, html, css } from 'lit-element';
import { styleMap } from "lit-html/directives/style-map";

export class Teammate extends LitElement {
    static get properties() {
        return {
            name: { type: String },
            status: { type: Number },
            selected: { type: Boolean },
        };
    }



    render() {
        let styles = {
            0: {
                color: '#000',
                style: 'normal'
            },
            1: {
                color: '#555',
                style: 'normal'
            },
            2: {
                color: '#555',
                style: 'italic'
            },
        }
        return html`
            <div style=${styleMap({
                color: styles[this.status].color,
                fontStyle: styles[this.status].style,
                padding: '0 5px',
                backgroundColor: this.selected ? '#ccc': 'inherit'
            })}>${this.name}</div>
        `;
    }
}
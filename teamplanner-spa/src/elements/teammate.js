import { LitElement, html, css } from 'lit-element';

export class Teammate extends LitElement {
    static get properties() {
        return {
            name: { type: String },
            status: { type: Number },
        };
    }

    render() {
        let styles = {
            0: {
                color: '#000',
                style: 'normal'
            },
            1: {
                color: '#aaa',
                style: 'normal'
            },
            2: {
                color: '#aaa',
                style: 'italic'
            },
        }
        return html`
            <div style="color: ${styles[this.status].color}; font-style: ${styles[this.status].style};">${this.name}</div>
        `;
    }
}
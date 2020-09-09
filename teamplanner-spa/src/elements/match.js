import { LitElement, html, css } from 'lit-element';

export class Match extends LitElement {
    static get properties() {
        return {
            date: { type: Date },
            description: { type: String },
        };
    }


    render() {
        return html`
            <div>${(new Date(this.date)).toLocaleDateString()}</div>
            <div style="font-size: small">${this.description}</div>
        `;
    }
}
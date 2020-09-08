import { LitElement, html, css } from 'lit-element';

export class Vote extends LitElement {
    static get properties() {
        return {
            vote: { type: Number },
            disabled: { type: Boolean },
        };
    }

    render() {
        let styles = {
            0: {
                glyph: "✔️",
                color: "#3fbd3f"
            },
            1: {
                glyph: "❌",
                color: "#bd3f3f"
            },
            2: {
                glyph: "❔",
                color: "#bdbd3f"
            },
        };
        return html`
            <div
                @click="${this._onClick}" 
                style="background-color: ${styles[this.vote].color}"
                >${styles[this.vote].glyph}</div>
        `;
    }

    _onClick() {
        if (this.disabled) 
            return;
        this.vote = (this.vote + 1) % 3
    }
}
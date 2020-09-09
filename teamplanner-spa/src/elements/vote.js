import { LitElement, html, css } from 'lit-element';

export class Vote extends LitElement {
    static get properties() {
        return {
            vote: { type: Number },
            disabled: { type: Boolean },
        };
    }

    static get styles() {
        return css`
        :host {
            -webkit-user-select: none;
                -ms-user-select: none;
                    user-select: none;
        }
        `;
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
            NaN: {
                glyph: "-",
                color: "#ccc"
            }
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

        // if the vote is still undefined, init it with 0 (i.e. yes)
        if (isNaN(this.vote)) {
            this.vote = 0
            return
        }

        // cycle through the three possible values (yes, no, maybe)
        this.vote = (this.vote + 1) % 3
    }
}
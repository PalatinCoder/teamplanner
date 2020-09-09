import { LitElement, html, css } from 'lit-element';
import { styleMap } from "lit-html/directives/style-map";

export class Vote extends LitElement {
    static get properties() {
        return {
            teammate: { type: Number },
            match: { type: String },
            vote: { type: Number },
            enabled: { type: Boolean },
            blur: { type: Boolean },
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
                style=${styleMap({
                    backgroundColor: styles[this.vote].color,
                    opacity: this.blur ? '0.4' : '1'
                })}
                >${styles[this.vote].glyph}</div>
        `;
    }

    _onClick() {
        console.log(`changing vote for teammate ${this.teammate} match ${this.match}`)
        if (!this.enabled) 
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
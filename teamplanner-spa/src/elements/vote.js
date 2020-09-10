import { LitElement, html, css } from 'lit-element';
import { styleMap } from "lit-html/directives/style-map";

export class Vote extends LitElement {
    static get properties() {
        return {
            vote: { type: Object },
            enabled: { type: Boolean },
            blur: { type: Boolean },
            _isOnline: { type: Boolean },
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

    constructor() {
        super();
        this._isOnline = navigator.onLine;
        window.addEventListener('online', () => this._isOnline = true);
        window.addEventListener('offline', () => this._isOnline = false);
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
            undefined: {
                glyph: "-",
                color: "#ccc"
            },
            999: {
                glyph: "⌛",
                color: "#ccc"
            }
        };
        return html`
            <div

                @click="${this._onClick}" 
                style=${styleMap({
                    backgroundColor: styles[this.vote.vote].color,
                    opacity: this.blur ? '0.4' : '1'
                })}
                >${styles[this.vote.vote].glyph}</div>
        `;
    }

    _onClick() {
        if (!(this.enabled && this._isOnline)) 
            return;

        // init the vote with 0, if it's undefined, otherwise cycle through the three possible states
        let newVote = isNaN(this.vote.vote) ? 0 : (this.vote.vote + 1) % 3

        this.vote.vote = 999 // loading indicator
        this.requestUpdate() // lit-element can't pick up that change deep in an object so we request an update manually

        this.dispatchEvent(new CustomEvent('vote-change', {
            detail: { teammate: this.vote.teammate, match: this.vote.match, vote: newVote },
            composed: true
        }))
    }
}
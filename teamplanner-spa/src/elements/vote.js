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
            NaN: {
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
                    backgroundColor: styles[this.vote].color,
                    opacity: this.blur ? '0.4' : '1'
                })}
                >${styles[this.vote].glyph}</div>
        `;
    }

    _onClick() {
        if (!(this.enabled && this._isOnline)) 
            return;

        // init the vote with 0, if it's undefined, otherwise cycle through the three possible states
        let oldVote = this.vote
        let newVote = isNaN(this.vote) ? 0 : (this.vote + 1) % 3

        this.vote = 999

        fetch('/vote', {
            method: 'POST',
            headers: { 'content-type': 'application/json' },
            body: JSON.stringify({ teammate: this.teammate, match: this.match, vote: newVote })
        })
        .then(() => { this.vote = newVote })
        .catch(error => { alert('Angabe konnte nicht gespeichert werden.'); console.log({error}); this.vote = oldVote })
    }
}
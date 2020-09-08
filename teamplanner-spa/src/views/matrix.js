import { LitElement, html, css } from 'lit-element';

export class VoteMatrix extends LitElement {
    static get properties() {
        return {
        }
    }

    static get styles() {
        return css`
        .wrapper {
            display: flex;
            width: 100%;
        }
        .list {
            display: flex;
            flex-direction: column;
            align-items: flex-start;
            margin-top: 55px;
            margin-right: 5px;
        }

        .grid {
            display: grid;
            grid-gap: 5px;
            grid-auto-columns: 100px;
            grid-template-rows: 50px auto;
            overflow-x: scroll;
        }

        tp-match {
            grid-row: 1;
        }

        tp-vote, tp-teammate {
            line-height: 25px;
        }

        tp-teammate {
            margin-bottom: 5px;
        }
        `;
    }

    render() {
        return html`
        <div class="wrapper">
            <div class="list">
                <tp-teammate name="Nat" status="0"></tp-teammate>
                <tp-teammate name="Torrey" status="1"></tp-teammate>
                <tp-teammate name="Heath" status="0"></tp-teammate>
                <tp-teammate name="Gaetano" status="2"></tp-teammate>
                <tp-teammate name="Nat" status="0"></tp-teammate>
                <tp-teammate name="Torrey" status="1"></tp-teammate>
                <tp-teammate name="Heath" status="0"></tp-teammate>
                <tp-teammate name="Gaetano" status="2"></tp-teammate>
            </div>
            <div class="grid">
                <tp-match date="02.01.2021" description="some text"></tp-match>
                <tp-vote vote="0" disabled></tp-vote>
                <tp-vote vote="1" disabled></tp-vote>
                <tp-vote vote="2" disabled></tp-vote>
                <tp-vote vote="2" disabled></tp-vote>
                <tp-vote vote="0" disabled></tp-vote>
                <tp-vote vote="1" disabled></tp-vote>
                <tp-vote vote="2" disabled></tp-vote>
                <tp-vote vote="2" disabled></tp-vote>
            </div>
        </div>
        
        `;
    }
}
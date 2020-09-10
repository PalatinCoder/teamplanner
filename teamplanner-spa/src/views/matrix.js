import { LitElement, html, css } from 'lit-element';

export class VoteMatrix extends LitElement {
    static get properties() {
        return {
            teammates: { type: Array },
            matches: { type: Array },
            votes: { type: Array },
            selectedTeammate: { type: Number },
        }
    }

    constructor() {
        super();
        this.selectedTeammate = 0;
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
            align-items: stretch;
            margin-top: 50px;
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
            margin-top: 5px;
        }
        `;
    }

    render() {
        return html`
        <div class="wrapper">
            <div class="list">
                ${this.teammates.map(
                    mate => html`<tp-teammate name="${mate.name}" status="${mate.status}" @click=${() => this._clickTeammate(mate.position)} ?selected=${this.selectedTeammate === mate.position}></tp-teammate>`
                )}
            </div>
            <div class="grid">
                ${this.matches.map(
                    match => html`<tp-match .date="${match.date}" .description="${match.description}"></tp-match>`
                )}

                ${this.generateVotesElements()}
            </div>
        </div>
        
        `;
    }

    _clickTeammate(pos) {
        if (this.selectedTeammate == pos) {
            this.selectedTeammate = 0 // cancel selection when clicked again
        } else {
            this.selectedTeammate = pos // set selection when clicked
        }
    }

    // Generate a vote box for each cell in the teammate x match matrix. It looks up every generated vote in the data from the API and sets the vote accordingly, or leaves it undefined
    generateVotesElements() {
        let elements = [] 
        for (var m = 0; m < this.matches.length; m++) {
            for (var t = 0; t < this.teammates.length; t++) {

                let teammate = this.teammates[t].position
                let match = this.matches[m].date.split("T")[0].replaceAll("-","") // holy shit this is ugly. "convert" JSON date to YYYYMMDD, as needed for the api
                let vote = this.votes.find(v => v.teammate == this.teammates[t].position && v.match == match) || { teammate: teammate.toString(), match, vote: undefined } // define a dummy object when no vote is found, so we can pass the property to the <tp-vote>

                let isSelected = this.selectedTeammate === teammate;
                let isNothingSelected = this.selectedTeammate == 0;
                let isAvailable = this.teammates[t].status === 0;
                let shouldNotBlur = isSelected || (isAvailable && isNothingSelected);

                // t starts at 0 while the grid-rows start at 1, so the top row being the matches the votes start at row 2
                // same thing for the columns, except there's no offset here
                let element = html`<tp-vote style="grid-row: ${t+2}; grid-column: ${m+1};"
                                        .vote="${vote}"
                                        ?enabled=${isSelected}
                                        ?blur=${!shouldNotBlur}>`

                elements = [...elements, element]
            }
        }
        return elements
    }
}
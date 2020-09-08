import { App } from './app.js';
import { Teammate } from './elements/teammate.js';
import { Match } from './elements/match.js';
import { Vote } from './elements/vote.js';
import { VoteMatrix } from './views/matrix.js';

customElements.define('tp-spa', App);

customElements.define('tp-teammate', Teammate)
customElements.define('tp-match', Match)
customElements.define('tp-vote', Vote)

customElements.define('vote-matrix', VoteMatrix);
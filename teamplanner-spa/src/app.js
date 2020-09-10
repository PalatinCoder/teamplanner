import { LitElement, html, css } from 'lit-element';
import moment from 'moment/dist/moment.js'
import PullToRefresh from "pulltorefreshjs";


export class App extends LitElement {
  static get properties() {
    return {
      teammates: { type: Array },
      matches: { type: Array },
      votes: { type: Array },
      isOffline: { type: Boolean },
    };
  }

  constructor() {
    super();
    this.teammates = [];
    this.matches = [];
    this.votes = [];
    this.isOffline = !navigator.onLine;

    window.addEventListener('online', () => { this.isOffline = false; this.fetchData() })
    window.addEventListener('offline', () => this.isOffline = true)
  }

  firstUpdated() {
    let that = this // oof... :D
    PullToRefresh.setPassiveMode(true)
    PullToRefresh.init({
      mainElement: 'body',
      onRefresh() { that.fetchData() }
    })

    this.fetchData()
  }

  async fetchData() {
    try {
      return Promise.all([
        fetch('/matches', { headers: { 'Accept': 'application/json' } })
          .then(r => r.json())
          .then(r_1 => r_1.filter(v => moment(v.date).isSameOrAfter(moment(), 'day')))
          .then(r_2 => { this.matches = r_2; }),
        fetch('/teammates', { headers: { 'Accept': 'application/json' } })
          .then(r_3 => r_3.json())
          .then(r_4 => { this.teammates = r_4; }),
        fetch('/votes', { headers: { 'Accept': 'application/json' } })
          .then(r_5 => r_5.json())
          .then(r_6 => { this.votes = r_6; })
      ]);
    }
    catch (error) {
      alert('Daten konnten nicht abgerufen werden.'); console.log(error);
    }
  }

  voteChange(event) {
    fetch('/vote', {
        method: 'POST',
        headers: { 'content-type': 'application/json' },
        body: JSON.stringify(event.detail)
    })
    .then (r => { if (!r.ok) throw Error(r.statusText)})
    .catch(error => { alert('Angabe konnte nicht gespeichert werden.'); console.log({error}); })
    // Could probably also store the new vote in memory instead of doing a full refresh *shrug*
    .finally(() => this.fetchData())
  }

  static get styles() {
    return css`
      :host {
        min-height: 100vh;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: flex-start;
        color: #1a2b42;
        max-width: 960px;
        margin: 0 auto;
        text-align: center;
      }

      header {
        display: flex;
        flex-wrap: wrap;
        align-items: center;
        width: 100%;
        margin: 1rem;
      }
      
      header img {
        max-width: 75px;
        max-height: 75px;
        margin: 0 1rem;
      }

      main {
        flex-grow: 1;
        width: 100%;
      }

      .app-footer {
        font-size: small;
        align-items: center;
      }

      #offline-notification {
        position: fixed;
        top: 0;
        width: 100%;
        padding: 5px;
        font-size: smaller;
        background-color: #aaa;
      }
    `;
  }

  render() {
    return html`
      ${this.isOffline ? html`
        <div id="offline-notification">🚫 Keine Internetverbindung<br>Zwischengespeicherte Daten, keine Änderungen möglich</div>
      ` : ''}
      <header>
        <img src="logo-default.png">
        <h1>Teamplaner</h1>
      </header>
      <main>
        <vote-matrix .teammates=${this.teammates} .matches=${this.matches} .votes=${this.votes} @vote-change=${this.voteChange}></vote-matrix>
      </main>

      <p class="app-footer">
        🏓 powered by
        <a
          target="_blank"
          rel="noopener noreferrer"
          href="https://jan-sl.de"
          >jan-sl.de</a
        >.
      </p>
    `;
  }
}

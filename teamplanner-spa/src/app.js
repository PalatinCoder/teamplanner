import { LitElement, html, css } from 'lit-element';
import moment from 'moment/dist/moment.js'


export class App extends LitElement {
  static get properties() {
    return {
      teammates: { type: Array },
      matches: { type: Array },
      votes: { type: Array },
      isOffline: { type: Boolean },
      isLoading: { type: Boolean },
    };
  }

  constructor() {
    super();
    this.teammates = [];
    this.matches = [];
    this.votes = [];
    this.isOffline = !navigator.onLine;
    this.isLoading = true;

    window.addEventListener('online', () => { this.isOffline = false; this.fetchData() })
    window.addEventListener('offline', () => this.isOffline = true)
  }

  firstUpdated() {
    this.fetchData()
  }

  fetchData() {
    this.isLoading = true;

    Promise.all([
    fetch('/matches', { headers: { 'Accept': 'application/json'}})
      .then(r => r.json())
      .then(r => r.filter(v => moment(v.date).isSameOrAfter(moment(), 'day')))
        .then(r => { this.matches = r }),
    fetch('/teammates', { headers: { 'Accept': 'application/json'}})
      .then(r => r.json())
        .then(r => { this.teammates = r }),
    fetch('/votes', { headers: { 'Accept': 'application/json'}})
    .then(r => r.json())
    .then(r => { this.votes = r })
    ]).catch(error => { alert('Daten konnten nicht abgerufen werden.'); console.log(error) })
    .finally(() => {
      this.isLoading = false
    })
    // TODO request an update
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

      #spinner {
        border: 16px solid #aaa;
        border-top: 16px solid #ccc;
        border-radius: 50%;
        width: 100px;
        height: 100px;
        animation: spin 2s linear infinite;
        position: absolute;
        top: 50%;
        left: 50%;
      }

      @keyframes spin {
        0% { transform: translate(-50%, -50%) rotate(0deg); }
        100% { transform: translate(-50%, -50%) rotate(360deg); }
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
        ${this.isLoading ? html`<div id="spinner"></div>` :''}
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

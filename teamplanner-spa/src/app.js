import { LitElement, html, css } from 'lit-element';

export class App extends LitElement {
  static get properties() {
    return {
      title: { type: String },
      page: { type: String },
    };
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
    `;
  }

  render() {
    return html`
      <header>
        <h1>Teamplanner</h1>
      </header>
      <main>
        <vote-matrix></vote-matrix>
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

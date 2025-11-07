  <h1>Go URL Shortener</h1>
  <p>A simple URL Shortener built using <strong>Go (Golang)</strong> with in-memory storage. 
  It generates short URLs using an MD5 hash and redirects users to the original URLs.</p>

  <hr />

  <h2>Features</h2>
  <ul>
    <li>Generate short URLs from long URLs</li>
    <li>Redirect short URLs to their original destinations</li>
    <li>Stores URLs in memory (no database required)</li>
    <li>Lightweight and fast Go HTTP server</li>
  </ul>

  <h2>Technologies Used</h2>
  <ul>
    <li>Go (net/http, encoding/json, crypto/md5)</li>
    <li>JSON API for communication</li>
  </ul>

  <h2> How It Works</h2>
  <ol>
    <li>User sends a POST request with the original URL to the <code>/shortner</code> endpoint.</li>
    <li>The server generates an 8-character MD5-based short URL.</li>
    <li>The short URL and original URL are stored in memory.</li>
    <li>When the user accesses <code>/redirect/{shortID}</code>, the server redirects them to the original URL.</li>
  </ol>

  <h2>API Endpoints</h2>

  <div class="endpoint">
    <h3>POST /shortner</h3>
    <p><strong>Description:</strong> Create a short URL for the given original URL.</p>
    <p><strong>Request Body:</strong></p>
    <pre><code>{
  "url": "https://example.com/very-long-url"
}</code></pre>

    <p><strong>Response:</strong></p>
    <pre><code>{
  "Short_url": "http://localhost:3000/redirect/a1b2c3d4"
}</code></pre>
  </div>

  <div class="endpoint">
    <h3>GET /redirect/{shortID}</h3>
    <p><strong>Description:</strong> Redirects to the original URL.</p>
    <p><strong>Example:</strong> <code>http://localhost:3000/redirect/a1b2c3d4</code></p>
  </div>

  <div class="endpoint">
    <h3>GET /</h3>
    <p><strong>Description:</strong> Root endpoint that displays a welcome message.</p>
  </div>

  <h2> Example Usage (cURL)</h2>

  <h3>1️Create Short URL</h3>
  <pre><code>curl -X POST http://localhost:3000/shortner \
-H "Content-Type: application/json" \
-d '{"url": "https://www.google.com"}'</code></pre>

  <h3>2️Redirect to Original URL</h3>
  <pre><code>http://localhost:3000/redirect/{shortID}</code></pre>


  <h2>Future Enhancements</h2>
  <ul>
    <li>Integrate persistent storage (SQLite, PostgreSQL, or Redis)</li>
    <li>Add analytics (click count, creation time, etc.)</li>
    <li>Add user authentication and custom short codes</li>
  </ul>

import { useState } from "react";
import countryData from "../data/countryOptions.js"; //i had some really weird issues here just importing my 
//...data and used AI for troubleshooting, never faced this one before. 
//Also this isnt a complete list of countries but it shows off the functionality of the app

export default function ValidateLogin() {
  //State components to hold IP address and Countries
  const [ipAddress, setIpAddress] = useState("");
  const [countries, setCountries] = useState([]);

// Fetch request to backend
const handleLogin = async () => {
  // this block was AI generated as i would not have added them on first go-round
  if (!ipAddress || countries.length === 0) {
    alert("Please enter both an IP address and select at least one country.");
    return;
  }

  try {
    const response = await fetch("http://localhost:8080/verify", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        ip: ipAddress,
        countries: countries,
      }),
    });

    const result = await response.json();

    // Ai suggestions helped form this statement
    if (response.ok) {
      alert(`Login ${result.allowed ? "Approved" : "Denied"}: ${result.message}`);
    }
  // AI helped with the catch block forming a better structure
  } catch (error) {
    console.error("Request failed:", error);
    alert("Request failed. Check console for details.");
  }
};
// AI helped me with the multi select dropdown which i had never done before and mapping the data from the JSON file. 
// it also suggested the "selected countries" display which was nice to have.
  return (
    <div className="login-container">
      <h2>Login Validation</h2>
      <input
        type="text"
        placeholder="IP Address"
        value={ipAddress}
        onChange={(e) => setIpAddress(e.target.value)}
      />
      <br/>
      <select
        multiple
        value={countries}
        onChange={(e) =>
          setCountries([...e.target.selectedOptions].map((o) => o.value))
        }
      >
        {countryData.map((country) => (
          <option key={country.alpha2} value={country.alpha2}>
            {country.name}
          </option>
        ))}
      </select>
      <br/>
      <p>Selected countries: {JSON.stringify(countries)}</p>
      <button onClick={handleLogin}>Validate</button>
    </div>
  );
}
# fi-mcp-dev

A minimal, hackathon-ready version of the Fi MCP server. This project provides a lightweight mock server for use in hackathons, demos, and development, simulating the core features of the production Fi MCP server with dummy data and simplified authentication.

## Purpose

- **fi-mcp-dev** is designed for hackathon participants and developers who want to experiment with the Fi MCP API without accessing real user data or production systems.
- It serves dummy financial data and uses a dummy authentication flow, making it safe and easy to use in non-production environments.

## Features

- **Simulates Fi MCP API**: Implements endpoints for net worth, credit report, EPF details, mutual fund transactions, and bank transactions.
- **Dummy Data**: All responses are served from static JSON files in `test_data_dir/`, representing various user scenarios.
- **Dummy Authentication**: Simple login flow using allowed phone numbers (directory names in `test_data_dir/`). No real OTP or user verification.
- **Hackathon-Ready**: No real integrations, no sensitive data, and easy to reset or extend.

## Directory Structure

- `main.go` — Entrypoint, sets up the server and endpoints.
- `middlewares/auth.go` — Implements dummy authentication and session management.
- `test_data_dir/` — Contains directories named after allowed phone numbers. Each directory holds JSON files for different API responses (e.g., `fetch_net_worth.json`).
- `static/` — HTML files for the login and login-successful pages.

## Dummy Data Scenarios

The dummy data covers a variety of user states. Example scenarios:

- **All assets connected**: Banks, EPF, Indian stocks, US stocks, credit report, large or small mutual fund portfolios.
- **All assets except bank account**: No bank account, but other assets present.
- **Multiple banks and UANs**: Multiple bank accounts and EPF UANs, partial transaction coverage.
- **No assets connected**: Only a savings account balance is present.
- **No credit report**: All assets except credit report.

## Test Data Scenarios

| Phone Number | Description                                                                                                                                                                                                                                        |
|-------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 1111111111  | No assets connected. Only saving account balance present                                                                                                                                                                                                                                            |
| 2222222222  | All assets connected (Banks account, EPF, Indian stocks, US stocks, Credit report). Large mutual fund portfolio with 9 funds                                                                                                                                                                        |
| 3333333333  | All assets connected (Banks account, EPF, Indian stocks, US stocks, Credit report). Small mutual fund portfolio with only 1 fund                                                                                                                                                                    |
| 4444444444  | All assets connected (Banks account, EPF, Indian stocks, US stocks). Small mutual fund portfolio with only 1 fund. With 2 UAN account connected . With 3 different bank with multiple account in them . Only have transactions for 2 bank accounts                                                  |
| 5555555555  | All assets connected except credit score (Banks account, EPF, Indian stocks, US stocks). Small mutual fund portfolio with only 1 fund. With 3 different bank with multiple account in them. Only have transactions for 2 bank accounts                                                              |
| 6666666666  | All assets connected except bank account (EPF, Indian stocks, US stocks). Large mutual fund portfolio with 9 funds. No bank account connected                                                                                                                                                       |
| 7777777777  | Debt-Heavy Low Performer. A user with mostly underperforming mutual funds, high liabilities (credit card & personal loans). Poor MF returns (XIRR < 5%). No diversification (all equity, few funds). Credit score < 650. High credit card usage, multiple loans. Negligible net worth or negative.  |
| 8888888888  | SIP Samurai. Consistently invests every month in multiple mutual funds via SIP. 3–5 active SIPs in MFs. Moderate returns (XIRR 8–12%).                                                                                                                                                              |
| 9999999999  | Fixed Income Fanatic. Strong preference for low-risk investments like debt mutual funds and fixed deposits. 80% of investments in debt MFs. Occasional gold ETF (Optional). Consistent but slow net worth growth (XIRR ~ 8-10%).                                                                    |
| 1010101010  | Precious Metal Believer. High allocation to gold and fixed deposits, minimal equity exposure. Gold MFs/ETFs ~50% of investment. Conservative SIPs in gold funds. FDs and recurring deposits. Minimal equity exposure.                                                                               |
| 1212121212  | Dormant EPF Earner. Has EPF account but employer stopped contributing; balance stagnant. EPF balance > ₹2 lakh. Interest not being credited. No private investment.                                                                                                                                 |
| 1414141414  | Salary Sinkhole. User’s salary is mostly consumed by EMIs and credit card bills. Salary credit every month. 70% goes to EMIs and credit card dues. Low or zero investment. Credit score ~600–650.                                                                                                   |
| 1313131313  | Balanced Growth Tracker. Well-diversified portfolio with EPF, MFs, stocks, and some US exposure. High EPF contribution. SIPs in equity & hybrid MFs. International MFs/ETFs 10–20%. Healthy net worth growth. Good credit score (750+).                                                             |
| 2020202020  | Starter Saver. Recently started investing, low ticket sizes, few transactions. Just 1–2 MFs, started < 6 months ago. SIP ₹500–₹1000. Minimal bank balance, no debt.                                                                                                                                 |
| 2121212121  | Dual Income Dynamo. Has freelance + salary income; cash flow is uneven but investing steadily. Salary + multiple credits from UPI apps. MF investments irregular but increasing. High liquidity in bank accounts. Credit score above 700. Occasional business loans or overdraft.                   |
| 2525252525  | Live-for-Today. High income but spends it all. Investments are negligible or erratic. Salary > ₹2L/month. High food, shopping, travel spends. No SIPs, maybe one-time MF buy. Credit card dues often roll over. Credit score < 700, low or zero net worth.                                          |

## Example: Dummy Data File

A sample `fetch_net_worth.json` (truncated for brevity):

```json
{
  "netWorthResponse": {
    "assetValues": [
      {"netWorthAttribute": "ASSET_TYPE_MUTUAL_FUND", "value": {"currencyCode": "INR", "units": "84642"}},
      {"netWorthAttribute": "ASSET_TYPE_EPF", "value": {"currencyCode": "INR", "units": "211111"}}
    ],
    "liabilityValues": [
      {"netWorthAttribute": "LIABILITY_TYPE_VEHICLE_LOAN", "value": {"currencyCode": "INR", "units": "5000"}}
    ],
    "totalNetWorthValue": {"currencyCode": "INR", "units": "658305"}
  }
}
```

## Authentication Flow

- When a tool/API is called, the server checks for a valid session.
- If not authenticated, the user is prompted to log in via a web page (`/mockWebPage?sessionId=...`).
- Enter any allowed phone number (see directories in `test_data_dir/`). OTP is not validated.
- On successful login, the session is stored in memory for the duration of the server run.

## Running the Server

### Prerequisites
- Go 1.23 or later ([installation instructions](https://go.dev/doc/install))

### Install dependencies
```sh
go mod tidy
```

### Start the server
```sh
FI_MCP_PORT=8080 go run .
```

The server will start on [http://localhost:8080](http://localhost:8080).

## Usage
- Follow instructions in this [guide](https://fi.money/features/getting-started-with-fi-mcp) to setup client
- Replace url with locally running server, for example: `http://localhost:8080/mcp/stream`
- When prompted for login, use one of the above phone numbers
- Otp/Passcode can be anything on the webpage

## Simple curl client
```bash
curl -X POST -H "Content-Type: application/json" -H "Mcp-Session-Id: mcp-session-594e48ea-fea1-40ef-8c52-7552dd9272af" -d '{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"fetch_bank_transactions","arguments":{}}}' http://localhost:8080/mcp/stream
```

If you run it once you will get `login_url` in response, running it again after login will give you the data

## Simple fastmcp python client
```py
from mcp.client.streamable_http import streamablehttp_client
from mcp.client.session import ClientSession
import asyncio
 
async def main():
    try:
        async with streamablehttp_client("http://localhost:8080/mcp/stream") as (
            read_stream,
            write_stream,
            _,
        ):
            async with ClientSession(
                read_stream,
                write_stream,
            ) as session:
                await session.initialize()
                tools = await session.list_tools()
                print(tools)
    except Exception as e:
        print(f"error: {e}")

if __name__ == "__main__":
    asyncio.run(main())

```

### Requirements:
- fastmcp (`pip install mcp`)

## Simple google adk example
```py
from google.adk.agents.llm_agent import LlmAgent
from google.adk.tools.mcp_tool.mcp_toolset import MCPToolset, StreamableHTTPConnectionParams

toolset = MCPToolset(
            connection_params=StreamableHTTPConnectionParams(
            url = "http://localhost:8080/mcp/stream"
            )
        )
root_agent = LlmAgent(
    model='gemini-2.0-flash',
    name='fi_mcp',
    instruction="You are an agent proficient in financial advises",
    tools=[toolset],
)
```
### Requirements:
- google-adk(>=1.8.0) (`pip install google-adk`)
- `export GOOGLE_API_KEY=YOUR_API_KEY`
- Run `adk web` in parent directory

If you are using `skip_auth` branch you should directly get data as shown
<img width="1134" height="545" alt="Screenshot 2025-07-27 at 03 57 08" src="https://github.com/user-attachments/assets/df4d377c-4bce-4d1a-b2da-48303b7907ea" />

## FAQ

- Why do I need to login everytime in ADKs?
  
  A session in MCP is one to one connection between MCP server and MCP client and needs login once. If your ADK have multiple agents and you are creating multiple clients for them make sure you maintain a common sessionId and pass that around. But in case you want to skip auth entirely, although not recommeneded because we want your agents to work on production usecases, you can refer to [this](https://github.com/epiFi/fi-mcp-dev/issues/3) issue on how to skip auth

- Why am I getting invalid session id?

  If you are creating custom session id before client initialization then make sure you prefix it with `mcp-session-`. For example: `3ef38b37-323a-4bbd-acbb-3fe02f97783f` is not valid and `mcp-session-3ef38b37-323a-4bbd-acbb-3fe02f97783f` is valid

defmodule LangxpayWeb.WelcomeController do
  use LangxpayWeb, :controller

  def index(conn, _params) do
    text(conn, "welcome to elixir")
  end
end
defmodule Langxpay.Users.Create do
  alias Langxpay.{Repo, User}

  def call(params) do
    params
    |> User.changeset()
    |> Repo.insert()
  end

end
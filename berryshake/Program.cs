using Microsoft.AspNetCore.Components.WebAssembly.Authentication;
using Microsoft.AspNetCore.Components.WebAssembly.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;
using MudBlazor.Services;

namespace Frontend
{
    public class Program
    {
        public static async Task Main(string[] args)
        {
            var gql_endpoint = "https://blue-surf-510082.us-east-1.aws.cloud.dgraph.io/graphql";
            var builder = WebAssemblyHostBuilder.CreateDefault(args);
            builder.RootComponents.Add<App>("#app");

            //builder.Services.AddScoped(sp => new HttpClient { BaseAddress = new Uri(builder.HostEnvironment.BaseAddress) });

            builder.Services.AddMudServices();

            // nuking the User-Agent header due to the CORS config of Dgraph Cloud. by default, a
            // custom User-Agent header is set by the `HttpClient`. this causes `user-agent` to be
            // added to the Access-Control-Request-Headers header in the OPTIONS request, but the
            // the server resp does not include it in Access-Control-Allow-Headers.
            // pattern taken from docs Strawberry Shake:
            //  https://chillicream.com/docs/strawberryshake/networking/authentication#httpclientfactory
            builder.Services.AddHttpClient(
                Frontend.GraphQL.ConferenceClient.ClientName,
                client => {
                    // Note: Dgraph Cloud will start breaking on CORS if a trailing slash is included
                    client.BaseAddress = new Uri(gql_endpoint);
                    client.DefaultRequestHeaders.Remove("User-Agent");
                })
                .AddHttpMessageHandler(sp => sp.GetRequiredService<AuthorizationMessageHandler>()
                    .ConfigureHandler(
                        authorizedUrls: new [] { gql_endpoint },
                        scopes: new[] { "openid",  "profile" }
                    )
                )
                .AddHttpMessageHandler(sp => sp.GetRequiredService<AuthHeaderOverwrite>());

            builder.Services.AddConferenceClient();
            builder.Services.AddScoped<AuthHeaderOverwrite>();

            // config taken from auth0 docs:
            //  https://auth0.com/blog/securing-blazor-webassembly-apps/
            builder.Services.AddOidcAuthentication(options =>
            {
                builder.Configuration.Bind("Auth0", options.ProviderOptions);
                options.ProviderOptions.ResponseType = "code";
            });

            await builder.Build().RunAsync();
        }
    }

    // more dumb CORS issues with Dgraph Cloud. the usual `Authorization` header isn't
    // whitelisted within `Access-Control-Allow-Headers`, so i grabbed one of the
    // random auth-related ones that are.
    class AuthHeaderOverwrite : DelegatingHandler
    {
        protected override Task<HttpResponseMessage> SendAsync(
            HttpRequestMessage request, System.Threading.CancellationToken cancellationToken)
        {
            var token = request.Headers.Authorization.ToString();
            request.Headers.Add("X-Auth-Token", token);
            request.Headers.Remove("Authorization");

            return base.SendAsync(request, cancellationToken);
        }
    }
}

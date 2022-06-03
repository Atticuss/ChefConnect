// <auto-generated/>
#pragma warning disable 1591
#pragma warning disable 0414
#pragma warning disable 0649
#pragma warning disable 0169

namespace frontend.Pages
{
    #line hidden
    using System;
    using System.Collections.Generic;
    using System.Linq;
    using System.Threading.Tasks;
    using Microsoft.AspNetCore.Components;
#nullable restore
#line 1 "/home/atticuss/github/chefconnect/frontend/_Imports.razor"
using System.Net.Http;

#line default
#line hidden
#nullable disable
#nullable restore
#line 2 "/home/atticuss/github/chefconnect/frontend/_Imports.razor"
using System.Net.Http.Json;

#line default
#line hidden
#nullable disable
#nullable restore
#line 3 "/home/atticuss/github/chefconnect/frontend/_Imports.razor"
using Microsoft.AspNetCore.Components.Forms;

#line default
#line hidden
#nullable disable
#nullable restore
#line 4 "/home/atticuss/github/chefconnect/frontend/_Imports.razor"
using Microsoft.AspNetCore.Components.Routing;

#line default
#line hidden
#nullable disable
#nullable restore
#line 5 "/home/atticuss/github/chefconnect/frontend/_Imports.razor"
using Microsoft.AspNetCore.Components.Web;

#line default
#line hidden
#nullable disable
#nullable restore
#line 6 "/home/atticuss/github/chefconnect/frontend/_Imports.razor"
using Microsoft.AspNetCore.Components.Web.Virtualization;

#line default
#line hidden
#nullable disable
#nullable restore
#line 7 "/home/atticuss/github/chefconnect/frontend/_Imports.razor"
using Microsoft.AspNetCore.Components.WebAssembly.Http;

#line default
#line hidden
#nullable disable
#nullable restore
#line 8 "/home/atticuss/github/chefconnect/frontend/_Imports.razor"
using Microsoft.JSInterop;

#line default
#line hidden
#nullable disable
#nullable restore
#line 9 "/home/atticuss/github/chefconnect/frontend/_Imports.razor"
using MudBlazor;

#line default
#line hidden
#nullable disable
#nullable restore
#line 10 "/home/atticuss/github/chefconnect/frontend/_Imports.razor"
using frontend;

#line default
#line hidden
#nullable disable
#nullable restore
#line 11 "/home/atticuss/github/chefconnect/frontend/_Imports.razor"
using frontend.Shared;

#line default
#line hidden
#nullable disable
#nullable restore
#line 3 "/home/atticuss/github/chefconnect/frontend/Pages/Ingredients.razor"
using frontend.Models;

#line default
#line hidden
#nullable disable
#nullable restore
#line 5 "/home/atticuss/github/chefconnect/frontend/Pages/Ingredients.razor"
using GraphQL;

#line default
#line hidden
#nullable disable
#nullable restore
#line 6 "/home/atticuss/github/chefconnect/frontend/Pages/Ingredients.razor"
using Microsoft.Extensions.Logging;

#line default
#line hidden
#nullable disable
    [Microsoft.AspNetCore.Components.RouteAttribute("/ingredients")]
    public partial class Ingredients : Microsoft.AspNetCore.Components.ComponentBase
    {
        #pragma warning disable 1998
        protected override void BuildRenderTree(Microsoft.AspNetCore.Components.Rendering.RenderTreeBuilder __builder)
        {
        }
        #pragma warning restore 1998
#nullable restore
#line 30 "/home/atticuss/github/chefconnect/frontend/Pages/Ingredients.razor"
       
    private List<Ingredient> ingredients;

    private async Task FetchIngredientData()
    {
        ingredients = await GraphqlOps.FetchIngredients();
        Logger.LogWarning("finished fetch");
        Logger.LogWarning("ingredients: " + ingredients);

        //GraphQL.GraphQLResponse<frontend.Models.ManyIngredients> gqlResp = await GraphqlOps.FetchIngredients();
        //Logger.LogWarning("gql: " + gqlResp);
        //Logger.LogWarning("gql.data: " + gqlResp.Data);
        //Logger.LogWarning("gql.data.ingredients: " + gqlResp.Data.ingredients);
        //ingredients = gqlResp.Data.ingredients;
    }

    // Executed immediately component is created
    protected override async Task OnInitializedAsync()
    {
        await FetchIngredientData();
    }

#line default
#line hidden
#nullable disable
        [global::Microsoft.AspNetCore.Components.InjectAttribute] private ILogger<Ingredients> Logger { get; set; }
        [global::Microsoft.AspNetCore.Components.InjectAttribute] private frontend.Data.GraphqlService GraphqlOps { get; set; }
    }
}
#pragma warning restore 1591

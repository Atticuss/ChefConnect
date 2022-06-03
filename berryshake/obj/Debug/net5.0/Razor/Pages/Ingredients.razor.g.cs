#pragma checksum "/home/atticuss/github/chefconnect/berryshake/Pages/Ingredients.razor" "{ff1816ec-aa5e-4d10-87f7-6f4963833460}" "aa6dae015d4a84719e1fac39ec4e51cb6415f43e"
// <auto-generated/>
#pragma warning disable 1591
namespace Frontend.Pages
{
    #line hidden
    using System;
    using System.Collections.Generic;
    using System.Linq;
    using System.Threading.Tasks;
    using Microsoft.AspNetCore.Components;
#nullable restore
#line 1 "/home/atticuss/github/chefconnect/berryshake/_Imports.razor"
using System.Net.Http;

#line default
#line hidden
#nullable disable
#nullable restore
#line 2 "/home/atticuss/github/chefconnect/berryshake/_Imports.razor"
using System.Net.Http.Json;

#line default
#line hidden
#nullable disable
#nullable restore
#line 3 "/home/atticuss/github/chefconnect/berryshake/_Imports.razor"
using Microsoft.AspNetCore.Authorization;

#line default
#line hidden
#nullable disable
#nullable restore
#line 4 "/home/atticuss/github/chefconnect/berryshake/_Imports.razor"
using Microsoft.AspNetCore.Components.Authorization;

#line default
#line hidden
#nullable disable
#nullable restore
#line 5 "/home/atticuss/github/chefconnect/berryshake/_Imports.razor"
using Microsoft.AspNetCore.Components.Forms;

#line default
#line hidden
#nullable disable
#nullable restore
#line 6 "/home/atticuss/github/chefconnect/berryshake/_Imports.razor"
using Microsoft.AspNetCore.Components.Routing;

#line default
#line hidden
#nullable disable
#nullable restore
#line 7 "/home/atticuss/github/chefconnect/berryshake/_Imports.razor"
using Microsoft.AspNetCore.Components.Web;

#line default
#line hidden
#nullable disable
#nullable restore
#line 8 "/home/atticuss/github/chefconnect/berryshake/_Imports.razor"
using Microsoft.AspNetCore.Components.Web.Virtualization;

#line default
#line hidden
#nullable disable
#nullable restore
#line 9 "/home/atticuss/github/chefconnect/berryshake/_Imports.razor"
using Microsoft.AspNetCore.Components.WebAssembly.Http;

#line default
#line hidden
#nullable disable
#nullable restore
#line 10 "/home/atticuss/github/chefconnect/berryshake/_Imports.razor"
using Microsoft.JSInterop;

#line default
#line hidden
#nullable disable
#nullable restore
#line 11 "/home/atticuss/github/chefconnect/berryshake/_Imports.razor"
using MudBlazor;

#line default
#line hidden
#nullable disable
#nullable restore
#line 12 "/home/atticuss/github/chefconnect/berryshake/_Imports.razor"
using Frontend;

#line default
#line hidden
#nullable disable
#nullable restore
#line 13 "/home/atticuss/github/chefconnect/berryshake/_Imports.razor"
using Frontend.Shared;

#line default
#line hidden
#nullable disable
#nullable restore
#line 15 "/home/atticuss/github/chefconnect/berryshake/_Imports.razor"
using Frontend.Models;

#line default
#line hidden
#nullable disable
#nullable restore
#line 4 "/home/atticuss/github/chefconnect/berryshake/Pages/Ingredients.razor"
using Frontend.GraphQL;

#line default
#line hidden
#nullable disable
#nullable restore
#line 5 "/home/atticuss/github/chefconnect/berryshake/Pages/Ingredients.razor"
using Microsoft.AspNetCore.Components.WebAssembly.Authentication;

#line default
#line hidden
#nullable disable
#nullable restore
#line 6 "/home/atticuss/github/chefconnect/berryshake/Pages/Ingredients.razor"
using System.ComponentModel;

#line default
#line hidden
#nullable disable
    [Microsoft.AspNetCore.Components.RouteAttribute("/ingredients")]
    public partial class Ingredients : Microsoft.AspNetCore.Components.ComponentBase
    {
        #pragma warning disable 1998
        protected override void BuildRenderTree(Microsoft.AspNetCore.Components.Rendering.RenderTreeBuilder __builder)
        {
            __builder.OpenComponent<MudBlazor.MudText>(0);
            __builder.AddAttribute(1, "Typo", global::Microsoft.AspNetCore.Components.CompilerServices.RuntimeHelpers.TypeCheck<MudBlazor.Typo>(
#nullable restore
#line 8 "/home/atticuss/github/chefconnect/berryshake/Pages/Ingredients.razor"
               Typo.h3

#line default
#line hidden
#nullable disable
            ));
            __builder.AddAttribute(2, "GutterBottom", global::Microsoft.AspNetCore.Components.CompilerServices.RuntimeHelpers.TypeCheck<System.Boolean>(
#nullable restore
#line 8 "/home/atticuss/github/chefconnect/berryshake/Pages/Ingredients.razor"
                                      true

#line default
#line hidden
#nullable disable
            ));
            __builder.AddAttribute(3, "ChildContent", (Microsoft.AspNetCore.Components.RenderFragment)((__builder2) => {
                __builder2.AddContent(4, "Ingredients");
            }
            ));
            __builder.CloseComponent();
            __builder.AddMarkupContent(5, "\n");
            __builder.OpenComponent<MudBlazor.MudText>(6);
            __builder.AddAttribute(7, "Class", "mb-8");
            __builder.AddAttribute(8, "ChildContent", (Microsoft.AspNetCore.Components.RenderFragment)((__builder2) => {
                __builder2.AddContent(9, "All currently available ingredients.");
            }
            ));
            __builder.CloseComponent();
            __builder.AddMarkupContent(10, "\n");
            __builder.OpenComponent<MudBlazor.MudButton>(11);
            __builder.AddAttribute(12, "Variant", global::Microsoft.AspNetCore.Components.CompilerServices.RuntimeHelpers.TypeCheck<MudBlazor.Variant>(
#nullable restore
#line 10 "/home/atticuss/github/chefconnect/berryshake/Pages/Ingredients.razor"
                    Variant.Text

#line default
#line hidden
#nullable disable
            ));
            __builder.AddAttribute(13, "Color", global::Microsoft.AspNetCore.Components.CompilerServices.RuntimeHelpers.TypeCheck<MudBlazor.Color>(
#nullable restore
#line 10 "/home/atticuss/github/chefconnect/berryshake/Pages/Ingredients.razor"
                                         Color.Primary

#line default
#line hidden
#nullable disable
            ));
            __builder.AddAttribute(14, "OnClick", global::Microsoft.AspNetCore.Components.CompilerServices.RuntimeHelpers.TypeCheck<Microsoft.AspNetCore.Components.EventCallback<Microsoft.AspNetCore.Components.Web.MouseEventArgs>>(Microsoft.AspNetCore.Components.EventCallback.Factory.Create<Microsoft.AspNetCore.Components.Web.MouseEventArgs>(this, 
#nullable restore
#line 10 "/home/atticuss/github/chefconnect/berryshake/Pages/Ingredients.razor"
                                                                 OpenNewIngredientDialog

#line default
#line hidden
#nullable disable
            )));
            __builder.AddAttribute(15, "ChildContent", (Microsoft.AspNetCore.Components.RenderFragment)((__builder2) => {
                __builder2.AddContent(16, "+ Add New");
            }
            ));
            __builder.CloseComponent();
#nullable restore
#line 12 "/home/atticuss/github/chefconnect/berryshake/Pages/Ingredients.razor"
 if (ingredients == null)
{

#line default
#line hidden
#nullable disable
            __builder.OpenComponent<MudBlazor.MudProgressCircular>(17);
            __builder.AddAttribute(18, "Color", global::Microsoft.AspNetCore.Components.CompilerServices.RuntimeHelpers.TypeCheck<MudBlazor.Color>(
#nullable restore
#line 14 "/home/atticuss/github/chefconnect/berryshake/Pages/Ingredients.razor"
                                Color.Default

#line default
#line hidden
#nullable disable
            ));
            __builder.AddAttribute(19, "Indeterminate", global::Microsoft.AspNetCore.Components.CompilerServices.RuntimeHelpers.TypeCheck<System.Boolean>(
#nullable restore
#line 14 "/home/atticuss/github/chefconnect/berryshake/Pages/Ingredients.razor"
                                                              true

#line default
#line hidden
#nullable disable
            ));
            __builder.CloseComponent();
#nullable restore
#line 15 "/home/atticuss/github/chefconnect/berryshake/Pages/Ingredients.razor"
}
else
{

#line default
#line hidden
#nullable disable
            __Blazor.Frontend.Pages.Ingredients.TypeInference.CreateMudTable_0(__builder, 20, 21, 
#nullable restore
#line 18 "/home/atticuss/github/chefconnect/berryshake/Pages/Ingredients.razor"
                     ingredients

#line default
#line hidden
#nullable disable
            , 22, 
#nullable restore
#line 18 "/home/atticuss/github/chefconnect/berryshake/Pages/Ingredients.razor"
                                         true

#line default
#line hidden
#nullable disable
            , 23, "Sort By", 24, 
#nullable restore
#line 18 "/home/atticuss/github/chefconnect/berryshake/Pages/Ingredients.razor"
                                                                              0

#line default
#line hidden
#nullable disable
            , 25, (__builder2) => {
                __builder2.OpenComponent<MudBlazor.MudTh>(26);
                __builder2.AddAttribute(27, "ChildContent", (Microsoft.AspNetCore.Components.RenderFragment)((__builder3) => {
                    __Blazor.Frontend.Pages.Ingredients.TypeInference.CreateMudTableSortLabel_1(__builder3, 28, 29, 
#nullable restore
#line 25 "/home/atticuss/github/chefconnect/berryshake/Pages/Ingredients.razor"
                                                 SortDirection.Ascending

#line default
#line hidden
#nullable disable
                    , 30, 
#nullable restore
#line 26 "/home/atticuss/github/chefconnect/berryshake/Pages/Ingredients.razor"
                        new Func<Ingredient, object>(x=>x.Name)

#line default
#line hidden
#nullable disable
                    , 31, (__builder4) => {
                        __builder4.AddContent(32, "Name");
                    }
                    );
                }
                ));
                __builder2.CloseComponent();
            }
            , 33, (context) => (__builder2) => {
                __builder2.OpenComponent<MudBlazor.MudTd>(34);
                __builder2.AddAttribute(35, "DataLabel", "Name");
                __builder2.AddAttribute(36, "ChildContent", (Microsoft.AspNetCore.Components.RenderFragment)((__builder3) => {
#nullable restore
#line 30 "/home/atticuss/github/chefconnect/berryshake/Pages/Ingredients.razor"
__builder3.AddContent(37, context.Name);

#line default
#line hidden
#nullable disable
                }
                ));
                __builder2.CloseComponent();
            }
            , 38, (__builder2) => {
                __builder2.OpenComponent<MudBlazor.MudTablePager>(39);
                __builder2.AddAttribute(40, "PageSizeOptions", global::Microsoft.AspNetCore.Components.CompilerServices.RuntimeHelpers.TypeCheck<System.Int32[]>(
#nullable restore
#line 33 "/home/atticuss/github/chefconnect/berryshake/Pages/Ingredients.razor"
                                        new int[]{50, 100}

#line default
#line hidden
#nullable disable
                ));
                __builder2.CloseComponent();
            }
            );
#nullable restore
#line 36 "/home/atticuss/github/chefconnect/berryshake/Pages/Ingredients.razor"
}

#line default
#line hidden
#nullable disable
        }
        #pragma warning restore 1998
#nullable restore
#line 38 "/home/atticuss/github/chefconnect/berryshake/Pages/Ingredients.razor"
       
    //private List<IGetIngredients_QueryIngredient> ingredients = new List<IGetIngredients_QueryIngredient>();
    public List<Ingredient> ingredients = new List<Ingredient>();
    protected override async Task OnInitializedAsync()
    {
        try
        {
            var result = await ConferenceClient.GetIngredients.ExecuteAsync();
            //ingredients = result.Data.QueryIngredient.ToList();
            ingredients = result.Data.QueryIngredient.Select(_ => new Ingredient{
                Id = _.Id,
                Name = _.Name,
            }).ToList();

        }
        catch (AccessTokenNotAvailableException exception)
        {
            exception.Redirect();
        }
    }

     async Task OpenNewIngredientDialog() {
         /*
         var parameters = new DialogParameters { ["ingredient"]=new GetIngredients_QueryIngredient_Ingredient("", "") };

        var dialog = DialogService.Show<IngredientDialog>("Add Ingredient", parameters);
        var result = await dialog.Result;

        if (!result.Cancelled)
        {
            Console.WriteLine("dialog completed successfully");
        } else {
            Console.WriteLine("dialog cancelled");
        }
        */
    }

#line default
#line hidden
#nullable disable
        [global::Microsoft.AspNetCore.Components.InjectAttribute] private IDialogService DialogService { get; set; }
        [global::Microsoft.AspNetCore.Components.InjectAttribute] private ConferenceClient ConferenceClient { get; set; }
    }
}
namespace __Blazor.Frontend.Pages.Ingredients
{
    #line hidden
    internal static class TypeInference
    {
        public static void CreateMudTable_0<T>(global::Microsoft.AspNetCore.Components.Rendering.RenderTreeBuilder __builder, int seq, int __seq0, global::System.Collections.Generic.IEnumerable<T> __arg0, int __seq1, global::System.Boolean __arg1, int __seq2, global::System.String __arg2, int __seq3, global::System.Int32 __arg3, int __seq4, global::Microsoft.AspNetCore.Components.RenderFragment __arg4, int __seq5, global::Microsoft.AspNetCore.Components.RenderFragment<T> __arg5, int __seq6, global::Microsoft.AspNetCore.Components.RenderFragment __arg6)
        {
        __builder.OpenComponent<global::MudBlazor.MudTable<T>>(seq);
        __builder.AddAttribute(__seq0, "Items", __arg0);
        __builder.AddAttribute(__seq1, "Hover", __arg1);
        __builder.AddAttribute(__seq2, "SortLabel", __arg2);
        __builder.AddAttribute(__seq3, "Elevation", __arg3);
        __builder.AddAttribute(__seq4, "HeaderContent", __arg4);
        __builder.AddAttribute(__seq5, "RowTemplate", __arg5);
        __builder.AddAttribute(__seq6, "PagerContent", __arg6);
        __builder.CloseComponent();
        }
        public static void CreateMudTableSortLabel_1<T>(global::Microsoft.AspNetCore.Components.Rendering.RenderTreeBuilder __builder, int seq, int __seq0, global::MudBlazor.SortDirection __arg0, int __seq1, global::System.Func<T, global::System.Object> __arg1, int __seq2, global::Microsoft.AspNetCore.Components.RenderFragment __arg2)
        {
        __builder.OpenComponent<global::MudBlazor.MudTableSortLabel<T>>(seq);
        __builder.AddAttribute(__seq0, "InitialDirection", __arg0);
        __builder.AddAttribute(__seq1, "SortBy", __arg1);
        __builder.AddAttribute(__seq2, "ChildContent", __arg2);
        __builder.CloseComponent();
        }
    }
}
#pragma warning restore 1591

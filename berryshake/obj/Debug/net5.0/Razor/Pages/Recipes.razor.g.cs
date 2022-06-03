#pragma checksum "/home/atticuss/github/chefconnect/berryshake/Pages/Recipes.razor" "{ff1816ec-aa5e-4d10-87f7-6f4963833460}" "e43ed5c4f5848d9f1c40ecfdb55a340fcb699dda"
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
#line 3 "/home/atticuss/github/chefconnect/berryshake/Pages/Recipes.razor"
using Frontend.GraphQL;

#line default
#line hidden
#nullable disable
#nullable restore
#line 4 "/home/atticuss/github/chefconnect/berryshake/Pages/Recipes.razor"
using Microsoft.AspNetCore.Components.WebAssembly.Authentication;

#line default
#line hidden
#nullable disable
#nullable restore
#line 5 "/home/atticuss/github/chefconnect/berryshake/Pages/Recipes.razor"
using System.ComponentModel;

#line default
#line hidden
#nullable disable
    [Microsoft.AspNetCore.Components.RouteAttribute("/recipes")]
    public partial class Recipes : Microsoft.AspNetCore.Components.ComponentBase
    {
        #pragma warning disable 1998
        protected override void BuildRenderTree(Microsoft.AspNetCore.Components.Rendering.RenderTreeBuilder __builder)
        {
            __builder.OpenComponent<MudBlazor.MudText>(0);
            __builder.AddAttribute(1, "Typo", global::Microsoft.AspNetCore.Components.CompilerServices.RuntimeHelpers.TypeCheck<MudBlazor.Typo>(
#nullable restore
#line 7 "/home/atticuss/github/chefconnect/berryshake/Pages/Recipes.razor"
               Typo.h3

#line default
#line hidden
#nullable disable
            ));
            __builder.AddAttribute(2, "GutterBottom", global::Microsoft.AspNetCore.Components.CompilerServices.RuntimeHelpers.TypeCheck<System.Boolean>(
#nullable restore
#line 7 "/home/atticuss/github/chefconnect/berryshake/Pages/Recipes.razor"
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
#nullable restore
#line 9 "/home/atticuss/github/chefconnect/berryshake/Pages/Recipes.razor"
 if (ingredients == null)
{

#line default
#line hidden
#nullable disable
            __builder.OpenComponent<MudBlazor.MudProgressCircular>(10);
            __builder.AddAttribute(11, "Color", global::Microsoft.AspNetCore.Components.CompilerServices.RuntimeHelpers.TypeCheck<MudBlazor.Color>(
#nullable restore
#line 11 "/home/atticuss/github/chefconnect/berryshake/Pages/Recipes.razor"
                                Color.Default

#line default
#line hidden
#nullable disable
            ));
            __builder.AddAttribute(12, "Indeterminate", global::Microsoft.AspNetCore.Components.CompilerServices.RuntimeHelpers.TypeCheck<System.Boolean>(
#nullable restore
#line 11 "/home/atticuss/github/chefconnect/berryshake/Pages/Recipes.razor"
                                                              true

#line default
#line hidden
#nullable disable
            ));
            __builder.CloseComponent();
#nullable restore
#line 12 "/home/atticuss/github/chefconnect/berryshake/Pages/Recipes.razor"
}
else
{

#line default
#line hidden
#nullable disable
            __Blazor.Frontend.Pages.Recipes.TypeInference.CreateMudTable_0(__builder, 13, 14, 
#nullable restore
#line 15 "/home/atticuss/github/chefconnect/berryshake/Pages/Recipes.razor"
                     ingredients

#line default
#line hidden
#nullable disable
            , 15, 
#nullable restore
#line 15 "/home/atticuss/github/chefconnect/berryshake/Pages/Recipes.razor"
                                         true

#line default
#line hidden
#nullable disable
            , 16, "Sort By", 17, 
#nullable restore
#line 15 "/home/atticuss/github/chefconnect/berryshake/Pages/Recipes.razor"
                                                                              0

#line default
#line hidden
#nullable disable
            , 18, (__builder2) => {
                __builder2.OpenComponent<MudBlazor.MudTh>(19);
                __builder2.AddAttribute(20, "ChildContent", (Microsoft.AspNetCore.Components.RenderFragment)((__builder3) => {
                    __Blazor.Frontend.Pages.Recipes.TypeInference.CreateMudTableSortLabel_1(__builder3, 21, 22, 
#nullable restore
#line 17 "/home/atticuss/github/chefconnect/berryshake/Pages/Recipes.razor"
                                                        SortDirection.Ascending

#line default
#line hidden
#nullable disable
                    , 23, 
#nullable restore
#line 17 "/home/atticuss/github/chefconnect/berryshake/Pages/Recipes.razor"
                                                                                         new Func<IGetIngredients_QueryIngredient, object>(x=>x.Name)

#line default
#line hidden
#nullable disable
                    , 24, (__builder4) => {
                        __builder4.AddContent(25, "Name");
                    }
                    );
                }
                ));
                __builder2.CloseComponent();
            }
            , 26, (context) => (__builder2) => {
                __builder2.OpenComponent<MudBlazor.MudTd>(27);
                __builder2.AddAttribute(28, "DataLabel", "Name");
                __builder2.AddAttribute(29, "ChildContent", (Microsoft.AspNetCore.Components.RenderFragment)((__builder3) => {
#nullable restore
#line 20 "/home/atticuss/github/chefconnect/berryshake/Pages/Recipes.razor"
__builder3.AddContent(30, context.Name);

#line default
#line hidden
#nullable disable
                }
                ));
                __builder2.CloseComponent();
            }
            , 31, (__builder2) => {
                __builder2.OpenComponent<MudBlazor.MudTablePager>(32);
                __builder2.AddAttribute(33, "PageSizeOptions", global::Microsoft.AspNetCore.Components.CompilerServices.RuntimeHelpers.TypeCheck<System.Int32[]>(
#nullable restore
#line 23 "/home/atticuss/github/chefconnect/berryshake/Pages/Recipes.razor"
                                            new int[]{50, 100}

#line default
#line hidden
#nullable disable
                ));
                __builder2.CloseComponent();
            }
            );
#nullable restore
#line 26 "/home/atticuss/github/chefconnect/berryshake/Pages/Recipes.razor"
}

#line default
#line hidden
#nullable disable
        }
        #pragma warning restore 1998
#nullable restore
#line 28 "/home/atticuss/github/chefconnect/berryshake/Pages/Recipes.razor"
       
    private List<IGetIngredients_QueryIngredient> ingredients = new List<IGetIngredients_QueryIngredient>();

    protected override async Task OnInitializedAsync()
    {
        try {
            var result = await ConferenceClient.GetIngredients.ExecuteAsync();
            ingredients = result.Data.QueryIngredient.ToList();
        
        } catch (AccessTokenNotAvailableException exception)
        {
            exception.Redirect();
        }
    }

#line default
#line hidden
#nullable disable
        [global::Microsoft.AspNetCore.Components.InjectAttribute] private ConferenceClient ConferenceClient { get; set; }
    }
}
namespace __Blazor.Frontend.Pages.Recipes
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
